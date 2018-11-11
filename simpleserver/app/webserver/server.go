// The Web server package.
// Provides http requests / response processing.

package webserver

import (
	"encoding/base64"
	"encoding/json"
	"github.com/karimarttila/go/simpleserver/app/domaindb"
	"github.com/karimarttila/go/simpleserver/app/userdb"
	"github.com/karimarttila/go/simpleserver/app/util"
	"log"
	"net/http"
	"strings"
)

type InfoMessage struct {
	Info string `json:"info"`
}

type SigninData struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Using ErrorResponder interface we can make error handling generic for all API calls.
type ErrorResponder interface {
	GetFlag() bool
	GetMsg() string
	WriteError(writer http.ResponseWriter) (err error)
}

// Used by all ErrorResponder entities.
func getEncoder(writer http.ResponseWriter) (encoder *json.Encoder) {
	encoder = json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	return encoder
}

// ErrorResponse is the base struct for all web layer Error response entities.
type ErrorResponse struct {
	Flag bool   `json:"-"` // Just to tell the whether we have initialized this struct or not (zero-value for bool is false, i.e. if the value is ready we know that we have initialized the struct).
	Ret  string `json:"ret"`
	Msg  string `json:"msg"`
}

func (e ErrorResponse) GetFlag() bool {
	return e.Flag
}

func (e ErrorResponse) GetMsg() string {
	return e.Msg
}

func (e ErrorResponse) WriteError(writer http.ResponseWriter) (err error) {
	encoder := getEncoder(writer)
	err = encoder.Encode(e)
	return err
}

// SigninErrorResponse is the /signin API error response entity.
type SigninErrorResponse struct {
	ErrorResponse
	Email string `json:"email"`
}

func (e SigninErrorResponse) GetFlag() bool {
	return e.Flag
}

func (e SigninErrorResponse) GetMsg() string {
	return e.Msg
}

func (e SigninErrorResponse) WriteError(writer http.ResponseWriter) (err error) {
	encoder := getEncoder(writer)
	err = encoder.Encode(e)
	return err
}

type SigninResponse struct {
	Flag  bool   `json:"-"`
	Ret   string `json:"ret"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Flag         bool   `json:"-"`
	Ret          string `json:"ret"`
	Msg          string `json:"msg"`
	JsonWebToken string `json:"json-web-token"`
}

// /info API.
func getInfo(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var httpStatus int
	infoMsg := &InfoMessage{Info: "index.html => Info in HTML format"}
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(infoMsg)
	if err != nil {
		util.LogError("JSON encoder returned error: " + err.Error())
		httpStatus = http.StatusBadRequest
	} else {
		httpStatus = http.StatusOK
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatus)
	util.LogExit()
}

func writeError(writer http.ResponseWriter, errorResponder ErrorResponder) {
	util.LogEnter()
	err := errorResponder.WriteError(writer)
	if err != nil {
		// Everything else failed, just write the json as string to http.ResponseWriter.
		writer.Write([]byte(`{"ret":"failed","msg":"A total failure, original error: ` + errorResponder.GetMsg() + `"}`))
	}
	util.LogExit()
}

func createErrorResponse(msg string) (errorResponse ErrorResponse) {
	util.LogEnter()
	ret := &ErrorResponse{true, "failed", msg}
	util.LogError(ret.GetMsg())
	errorResponse = *ret
	util.LogExit()
	return errorResponse
}

// TODO: it would be nice to make this generic as well.
func createSigninErrorResponse(msg string, email string) (signinErrorResponse SigninErrorResponse) {
	util.LogEnter()
	ret := &SigninErrorResponse{
		ErrorResponse: ErrorResponse{true, "failed", msg},
		Email:         email,
	}
	util.LogError(ret.GetMsg())
	signinErrorResponse = *ret
	util.LogExit()
	return signinErrorResponse
}

func writeHeaders(writer http.ResponseWriter, errorResponder ErrorResponder) {
	writer.Header().Set("Content-Type", "application/json")
	if errorResponder.GetFlag() {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

func postSignin(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var signinErrorResponse SigninErrorResponse
	var signinData SigninData
	var signinResponse SigninResponse
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&signinData)
	if err != nil {
		signinErrorResponse = createSigninErrorResponse("Decoding request body failed", "")
	} else {
		if signinData.FirstName == "" || signinData.LastName == "" || signinData.Email == "" || signinData.Password == "" {
			signinErrorResponse = createSigninErrorResponse("Validation failed - some fields were empty", "")
		} else {
			var ret userdb.AddUserResponse
			ret, err = userdb.AddUser(signinData.Email, signinData.FirstName, signinData.LastName, signinData.Password)
			if err != nil {
				signinErrorResponse = createSigninErrorResponse(err.Error(), signinData.Email)
			} else {
				util.LogTrace("AddUser returned: Ret: " + ret.Ret + ", Email: " + ret.Email)
				signinResponse = SigninResponse{true, "ok", signinData.Email}
				encoder := json.NewEncoder(writer)
				encoder.SetEscapeHTML(false)
				err := encoder.Encode(signinResponse)
				if err != nil {
					signinErrorResponse = createSigninErrorResponse(err.Error(), signinData.Email)
				}
			}
		}
	}
	writeHeaders(writer, signinErrorResponse)
	if signinErrorResponse.Flag {
		writeError(writer, signinErrorResponse)
	}
	util.LogExit()
}

func postLogin(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var errorResponse ErrorResponse // Generic ErrorResponse will do for /login just fine.
	var loginData LoginData
	var loginResponse LoginResponse
	var jsonWebToken string
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&loginData)
	if err != nil {
		errorResponse = createErrorResponse("Decoding request body failed")
	} else {
		if loginData.Email == "" || loginData.Password == "" {
			errorResponse = createErrorResponse("Validation failed - some fields were empty")
		} else {
			credentialsOk := userdb.CheckCredentials(loginData.Email, loginData.Password)
			if !credentialsOk {
				errorResponse = createErrorResponse("Credentials are not good - either email or password is not correct")
			} else {
				jsonWebToken, err = CreateJsonWebToken(loginData.Email)
				if err != nil {
					errorResponse = createErrorResponse("Couldn't create token: " + err.Error())
				} else {
					loginResponse = LoginResponse{true, "ok", "Credentials ok", jsonWebToken}
					encoder := json.NewEncoder(writer)
					encoder.SetEscapeHTML(false)
					err := encoder.Encode(loginResponse)
					if err != nil {
						errorResponse = createErrorResponse(err.Error())
					}
				}
			}
		}
	}
	writeHeaders(writer, errorResponse)
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
	util.LogExit()
}

func isValidToken(request *http.Request) (email string, errorResponse ErrorResponse) {
	util.LogEnter()
	auth := request.Header.Get("Authorization")
	if auth == "" {
		errorResponse = createErrorResponse("Authorization not found in the header parameters")
	} else {
		util.LogTrace("Got auth: " + auth)
		authRest := auth[6:] // Get rid of "Basic "
		decodedBytes, err := base64.StdEncoding.DecodeString(authRest)
		if err != nil {
			errorResponse = createErrorResponse("Couldn't base64 decode auth string: " + err.Error())
		} else {
			decoded := string(decodedBytes)
			util.LogTrace("decoded: " + decoded)
			index := strings.Index(decoded, ":NOT")
			var token string
			if index == -1 {
				token = decoded
			} else {
				token = decoded[0:index]
			}
			util.LogTrace("token: " + token)
			var tokenResponse TokenResponse
			tokenResponse, err = ValidateJsonWebToken(token)
			if err != nil {
				errorResponse = createErrorResponse("Couldn't validate token: " + err.Error())
			} else {
				util.LogTrace("tokenResponse.email: " + tokenResponse.Email)
				email = tokenResponse.Email
			}
		}
	}
	util.LogExit()
	return email, errorResponse
}


func getProductGroups(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	parsedEmail, errorResponse := isValidToken(request)
	var productGroups domaindb.ProductGroups
	if !errorResponse.Flag {
		util.LogTrace("parsedEmail from token: " + parsedEmail)
		productGroups = domaindb.GetProductGroups()
		encoder := json.NewEncoder(writer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(productGroups)
		if err != nil {
			errorResponse = createErrorResponse(err.Error())
		}
	}
	writeHeaders(writer, errorResponse)
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
	util.LogExit()
}

// Registers the API calls.
func handleRequests() {
	util.LogEnter()
	http.HandleFunc("/info", getInfo)
	http.HandleFunc("/signin", postSignin)
	http.HandleFunc("/login", postLogin)
	http.HandleFunc("/product-groups", getProductGroups)
	http.Handle("/", http.FileServer(http.Dir("./src/github.com/karimarttila/go/simpleserver/static")))
	log.Fatal(http.ListenAndServe(":"+util.MyConfig["port"], nil))
	util.LogExit()
}

// The main entry point to the file.
// Remember that exportable functions begin with a capital letter.
func StartServer() {
	util.LogEnter()
	handleRequests()
	util.LogExit()
}

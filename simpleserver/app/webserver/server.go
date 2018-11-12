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
	"strconv"
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


// /info API.
func getInfo(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var errorResponse ErrorResponse // Generic ErrorResponse will do for /info just fine.
	if !(request.Method == "GET") {
		errorResponse = createErrorResponse("Only GET allowed for /info")
	} else {
		infoMsg := &InfoMessage{Info: "index.html => Info in HTML format"}
		encoder := json.NewEncoder(writer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(infoMsg)
		if err != nil {
			errorResponse = createErrorResponse("JSON encoder returned error: " + err.Error())
		}
	}
	writeHeaders(writer, errorResponse)
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
	util.LogExit()
}


func postSignin(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var signinErrorResponse SigninErrorResponse
	var signinData SigninData
	var signinResponse SigninResponse
	if !(request.Method == "POST") {
		signinErrorResponse = createSigninErrorResponse("Only POST allowed for /signin", "")
	} else {
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
	if !(request.Method == "POST") {
		errorResponse = createErrorResponse("Only POST allowed for /login")
	} else {
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
	if !(request.Method == "GET") {
		errorResponse = createErrorResponse("Only GET allowed for /product-groups")
	} else {
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
	}
	writeHeaders(writer, errorResponse)
	if errorResponse.Flag {
		writeError(writer, errorResponse)
	}
	util.LogExit()
}

func getProducts(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var parsedEmail string
	var errorResponse ErrorResponse
	var pgId int
	var err error
	var products domaindb.Products
	if !(request.Method == "GET") {
		errorResponse = createErrorResponse("Only GET allowed for /products")
	} else {
		parsedEmail, errorResponse = isValidToken(request)
		util.LogTrace("parsedEmail: " + parsedEmail)
		if !errorResponse.Flag {
			// like: /products/1
			pgIdStr := request.URL.Path[len("/products/"):]
			if len(pgIdStr) < 1 {
				errorResponse = createErrorResponse("pgId was less than 1")
			} else {
				pgId, err = strconv.Atoi(pgIdStr)
				if err != nil {
					errorResponse = createErrorResponse("pgId was not an integer")
				} else {
					util.LogTrace("pgId: " + string(pgId))
					products = domaindb.GetProducts(pgId)
					encoder := json.NewEncoder(writer)
					encoder.SetEscapeHTML(false)
					err := encoder.Encode(products)
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


func getProduct(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var parsedEmail string
	var errorResponse ErrorResponse
	var pgId, pId int
	var err error
	var rawProduct domaindb.RawProduct
	if !(request.Method == "GET") {
		errorResponse = createErrorResponse("Only GET allowed for /product")
	} else {
		parsedEmail, errorResponse = isValidToken(request)
		util.LogTrace("parsedEmail: " + parsedEmail)
		if !errorResponse.Flag {
			// like: /product/1
			idsStr := request.URL.Path[len("/product/"):]
			if len(idsStr) < 1 {
				errorResponse = createErrorResponse("idsStr was less than 1")
			} else {
				ids := strings.Split(idsStr, "/")
				if len(ids) != 2 {
					errorResponse = createErrorResponse("We didn't find both product group id and product id in the url parameters")
				} else {
					pgId, err = strconv.Atoi(ids[0])
					if err != nil {
						errorResponse = createErrorResponse("pgId was not an integer")
					} else {
						pId, err = strconv.Atoi(ids[1])
						if err != nil {
							errorResponse = createErrorResponse("pId was not an integer")
						} else {
							util.LogTrace("pgId: " + string(pgId) + ", pId: " + string(pId))
							rawProduct = domaindb.GetProduct(pgId, pId)
							encoder := json.NewEncoder(writer)
							encoder.SetEscapeHTML(false)
							err := encoder.Encode(rawProduct)
							if err != nil {
								errorResponse = createErrorResponse(err.Error())
							}
						}
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


// Registers the API calls.
func handleRequests() {
	util.LogEnter()
	http.HandleFunc("/info", getInfo)
	http.HandleFunc("/signin", postSignin)
	http.HandleFunc("/login", postLogin)
	http.HandleFunc("/product-groups", getProductGroups)
	http.HandleFunc("/products/", getProducts)
	http.HandleFunc("/product/", getProducts)
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

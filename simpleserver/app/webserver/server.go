// The Web server package.
// Provides http requests / response processing.

package webserver

import (
	"encoding/json"
	"github.com/karimarttila/go/simpleserver/app/userdb"
	"github.com/karimarttila/go/simpleserver/app/util"
	"log"
	"net/http"
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

type ErrorResponse struct {
	Flag bool // Just to tell the whether we have initialized this struct or not (zero-value for bool is false, i.e. if the value is ready we know that we have initialized the struct).
	Ret  string `json:"ret"`
	Msg  string `json:"msg"`
}

type SigninResponse struct {
	Flag  bool
	Ret   string `json:"ret"`
	Email string `json:"email"`
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

func errorHandler(msg string) (errorResponse ErrorResponse, httpStatus int) {
	util.LogEnter()
	util.LogError(msg)
	errorResponse = ErrorResponse{true, "failed", msg}
	httpStatus = http.StatusBadRequest
	util.LogExit()
	return errorResponse, httpStatus
}

// Last resort error handler.
func writeError(writer http.ResponseWriter, errorRet ErrorResponse) {
	util.LogEnter()
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(errorRet)
	if err != nil {
		// Everything else failed, just write the json as string to http.ResponseWriter.
		writer.Write([]byte(`{"ret":"failed","msg":"A total failure, original error: ` + errorRet.Msg + `"}`))
	}
	util.LogExit()
}

func postSignin(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	var errorRet ErrorResponse
	var signinData SigninData
	var httpStatus int
	var signinResponse SigninResponse
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&signinData)
	if err != nil {
		errorRet, httpStatus = errorHandler("Decoding request body failed")
	} else {
		var ret userdb.AddUserResponse
		ret, err = userdb.AddUser(signinData.Email, signinData.FirstName, signinData.LastName, signinData.Password)
		if err != nil {
			errorRet, httpStatus = errorHandler(err.Error())
		} else {
			signinResponse = SigninResponse{true, "ok", ret.Email}
			encoder := json.NewEncoder(writer)
			encoder.SetEscapeHTML(false)
			err := encoder.Encode(signinResponse)
			if err != nil {
				errorRet, httpStatus = errorHandler(err.Error())
			}
		}
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatus)
	if err != nil {
		writeError(writer, errorRet)
	}
	util.LogExit()
}

// Registers the API calls.
func handleRequests() {
	util.LogEnter()
	http.HandleFunc("/info", getInfo)
	http.HandleFunc("/signin", postSignin)
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

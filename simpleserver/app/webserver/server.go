// The Web server package.
// Provides http requests / response processing.

package webserver

import (
	"bytes"
	"encoding/json"
	"github.com/karimarttila/go/simpleserver/app/util"
	"log"
	"net/http"
)

type InfoMessage struct {
	Info string `json:"info"`
}

//
func JSONMarshalPreserveHTMLCharacters(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// /info API.
func getInfo(writer http.ResponseWriter, request *http.Request) {
	util.LogEnter()
	infoMsg := &InfoMessage{Info: "index.html => Info in HTML format"}
	retBytes, err := JSONMarshalPreserveHTMLCharacters(infoMsg)
	httpStatus := http.StatusOK
	if err != nil {
		retBytes = []byte(`{"info":"JSON parsing failed in getInfo"}`)
		httpStatus = http.StatusBadRequest
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatus)
	writer.Write(retBytes)
	util.LogExit()
}

//{"info":"index.html => Info in HTML format"}

// Registers the API calls.
func handleRequests() {
	util.LogEnter()
	http.HandleFunc("/info", getInfo)
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

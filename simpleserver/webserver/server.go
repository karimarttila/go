// The Web server package.
// Provides http requests / response processing.

package webserver

import (
	"fmt"
	"github.com/karimarttila/go/simpleserver/util"
	"log"
	"net/http"
	"strconv"
)


// /info API.
func getInfo(w http.ResponseWriter, r *http.Request) {
	util.LogEnter()
	fmt.Fprintf(w, "TODO: This is info")
	util.LogExit()
}


// Registers the API calls.
func handleRequests() {
	util.LogEnter()
	http.HandleFunc("/info", getInfo)
	http.Handle("/", http.FileServer(http.Dir("./src/github.com/karimarttila/go/simpleserver/static")))
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(util.MyConfig.Port), nil))
	util.LogExit()
}

// The main entry point to the file.
// Remember that exportable functions begin with a capital letter.
func StartServer() {
	util.LogEnter()
	handleRequests()
	util.LogExit()
}


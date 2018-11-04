// The Web server package.
// Provides http requests / response processing.

package webserver

import (
	"fmt"
	"github.com/karimarttila/go/simpleserver/util"
	"log"
	"net/http"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	util.MyLogger.Debug(util.ENTER)
	fmt.Fprintf(w, "This is info")
	util.MyLogger.Debug(util.EXIT)
}

func handleRequests() {
	util.MyLogger.Debug(util.ENTER)
	http.HandleFunc("/info", getInfo)
	http.Handle("/", http.FileServer(http.Dir("./src/github.com/karimarttila/go/simpleserver/static")))
	log.Fatal(http.ListenAndServe(":" + util.MyConfig.Port, nil))
	util.MyLogger.Debug(util.EXIT)
}

// The main entry point to the file.
// Remember that exportable functions begin with a capital letter.
func StartServer() {
	util.MyLogger.Debug(util.ENTER)
	handleRequests()
	util.MyLogger.Debug(util.EXIT)
}

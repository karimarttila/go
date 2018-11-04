// The Web server package.
// Provides http requests / response processing.

package webserver

import (
	"fmt"
	"log"
	"net/http"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is info")
}

func handleRequests() {
	http.HandleFunc("/info", getInfo)

	http.Handle("/", http.FileServer(http.Dir("./src/github.com/karimarttila/go/simpleserver/static")))

	log.Fatal(http.ListenAndServe(":3048", nil))
}

// The main entry point to the file.
// Remember that exportable functions begin with a capital letter.
func StartServer() {
	handleRequests()
}

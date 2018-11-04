// The main entry point to the Simple Server Go application.
package main

import (
	"github.com/karimarttila/go/simpleserver/webserver"
)

// The main entry point to the file.
// Just calls the webserver package to start the http server.
func main() {
	webserver.StartServer()
}
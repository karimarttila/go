// The main entry point to the Simple Server Go application.
package main

import (
	"github.com/karimarttila/go/simpleserver/util"
	"github.com/karimarttila/go/simpleserver/webserver"
)


// The main entry point to the file.
// Just calls the webserver package to start the http server.
func main() {
	util.MyLogger.Debug(util.ENTER)
	util.MyLogger.Debug("Starting server...")
	util.MyLogger.Debug("- port: " + util.MyConfig.Port)

	webserver.StartServer()
	util.MyLogger.Debug(util.ENTER)
}

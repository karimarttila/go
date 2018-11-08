// The main entry point to the Simple Server Go application.
package main

import (
	"github.com/karimarttila/go/simpleserver/app/util"
	"github.com/karimarttila/go/simpleserver/app/webserver"
)

// The main entry point to the file.
// Just calls the webserver package to start the http server.
func main() {
	util.LogEnter()
	util.LogDebug("Starting server...")
	util.LogDebug("- port: " + util.MyConfig["port"])
	util.LogDebug("- report_caller: " + util.MyConfig["report_caller"])
	util.LogDebug("- log_level: " + util.MyConfig["log_level"])
	util.LogDebug("- log_file: " + util.MyConfig["log_file"])
	webserver.StartServer()
	util.LogExit()
	// Finally close the log file.
	util.CloseLog()
}

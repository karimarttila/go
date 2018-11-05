// The main entry point to the Simple Server Go application.
package main

import (
	"github.com/karimarttila/go/simpleserver/util"
	"github.com/karimarttila/go/simpleserver/webserver"
	"strconv"
)


// The main entry point to the file.
// Just calls the webserver package to start the http server.
func main() {
	util.LogEnter()
	util.LogDebug("Starting server...")
	util.LogDebug("- Port: " + strconv.Itoa(util.MyConfig.Port))
	util.LogDebug("- Report_caller: " + strconv.FormatBool(util.MyConfig.Report_caller))
	util.LogDebug("- Log_level: " + util.MyConfig.Log_level)
	util.LogDebug("- Log_file: " + util.MyConfig.Log_file)
	webserver.StartServer()
	util.LogExit()
}

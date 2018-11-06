package util

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// TODO: Make your own implementation based on Go standard library and not use extra dependencies!

var MyConfig = initConfiguration()

type Configuration struct {
	Port              int
	Report_caller     bool
	Log_level         string
	Log_file          string
}

// Initializes the configuration.
func initConfiguration() (ret Configuration){
	// NOTE: We cannot use the logger yet since it uses MyConfig which is not yet initialized (or initialization loop happens).
	fmt.Println("simpleserver.util.config.go - initConfiguration - ENTER")
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println("simpleserver.util.config.go - initConfiguration - ERROR: " + err.Error())
		os.Exit(500)
	}
	fmt.Println("simpleserver.util.config.go - initConfiguration - EXIT")
	return configuration
}


func getFileName() string {
	fmt.Println("simpleserver.util.config.go - getFileName - ENTER")
	env := os.Getenv("SS_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"../config", "/config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	fmt.Println("simpleserver.util.config.go - getFileName - EXIT")
	return filePath
}

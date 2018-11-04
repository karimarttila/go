package util

import (
	"github.com/tkanos/gonfig"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var MyConfig = initConfiguration()

type Configuration struct {
	Port              string
	Connection_String string
}

func initConfiguration() (ret Configuration){
	MyLogger.Debug(ENTER)
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		MyLogger.Error(err)
		os.Exit(500)
	}
	MyLogger.Debug(EXIT)
	return configuration
}


func getFileName() string {
	MyLogger.Debug(ENTER)
	env := os.Getenv("SS_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"../config", "/config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	MyLogger.Debug(EXIT)
	return filePath
}

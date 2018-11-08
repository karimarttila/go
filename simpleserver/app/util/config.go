package util

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// A simple map comprising the Simple Server properties.
var MyConfig = readProperties()

type Config map[string]string

func readProperties() Config {
	fmt.Println("simpleserver.util.config. - readProperties - ENTER")
	config := Config{}
	filename := getFileName()
	if len(filename) == 0 {
		fmt.Println("simpleserver.util.config.go - readProperties - ERROR: file name was empty")
		os.Exit(500)
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("simpleserver.util.config.go - readProperties - ERROR: error opening file: " + err.Error())
		os.Exit(500)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("simpleserver.util.config.go - readProperties - ERROR: error while scanning file: " + err.Error())
		os.Exit(500)
	}
	fmt.Println("simpleserver.util.config. - readProperties - EXIT")
	return config
}

func getFileName() string {
	fmt.Println("simpleserver.util.config.go - getFileName - ENTER")
	env := os.Getenv("SS_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"../../config", "/" + env + "-config.properties"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	fmt.Println("simpleserver.util.config.go - getFileName - EXIT")
	return filePath
}

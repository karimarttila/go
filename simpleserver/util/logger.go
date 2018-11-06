// Logging configuration package.
// NOTE: In production code report caller (Report_caller) is expensive and should be turned off.
// Provides two helper methods for logging function entry and exit.
package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	log.SetFlags(0)
	file, err := os.OpenFile(MyConfig.Log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	} else {
		log.Fatal("Failed to open log file %s, using just stdout", MyConfig.Log_file)
		log.SetOutput(os.Stdout)
	}
}

type SSLogLevel int

const (
	SS_LOG_LEVEL_TRACE SSLogLevel = iota
	SS_LOG_LEVEL_DEBUG
	SS_LOG_LEVEL_INFO
	SS_LOG_LEVEL_WARN
	SS_LOG_LEVEL_ERROR
	SS_LOG_LEVEL_FATAL
)

var MyLogLevel = initLogLevel()

func initLogLevel() (SSLogLevel) {
	var myLogLevel SSLogLevel
	switch MyConfig.Log_level {
	case "trace":
		myLogLevel = SS_LOG_LEVEL_TRACE
	case "debug":
		myLogLevel = SS_LOG_LEVEL_DEBUG
	case "info":
		myLogLevel = SS_LOG_LEVEL_INFO
	case "warn":
		myLogLevel = SS_LOG_LEVEL_WARN
	case "error":
		myLogLevel = SS_LOG_LEVEL_ERROR
	case "fatal":
		myLogLevel = SS_LOG_LEVEL_FATAL
	default:
		fmt.Println("simpleserver.util.logger.go - getLogLevel - ERROR: Unknown log level: " + MyConfig.Log_level)
		os.Exit(500)
	}
	return myLogLevel
}

func (level SSLogLevel) String() string {
	levels := [...]string {
		"TRACE",
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL"}
	if level < SS_LOG_LEVEL_TRACE || level > SS_LOG_LEVEL_FATAL {
		return "Unknown SS_LOG_LEVEL"
	}
	ret := levels[level]
	return ret
}

func logIt(msg string, level SSLogLevel) {
	var caller string
	var entry string
	var timeStamp = fmt.Sprintf(time.Now().UTC().Format("2006-01-02T15:04:05.999Z"))
	if level >= MyLogLevel {
		if MyConfig.Report_caller {
			pc, _, _, _ := runtime.Caller(2)
			fn := runtime.FuncForPC(pc)
			caller = fn.Name()
			caller = strings.Replace(caller, "github.com/karimarttila/go/simpleserver/", "", 1)
			entry = fmt.Sprintf("[%s] - [%s] [%s] - %s", timeStamp, level, caller, msg)
		} else {
			entry = fmt.Sprintf("[%s] - [%s] - %s", timeStamp, level, msg)
		}
		log.Println(entry)
	}
}

func LogTrace(msg string) {
	logIt(msg, SS_LOG_LEVEL_TRACE)
}

func LogDebug(msg string) {
	logIt(msg, SS_LOG_LEVEL_DEBUG)
}

func LogInfo(msg string) {
	logIt(msg, SS_LOG_LEVEL_INFO)
}

func LogWarn(msg string) {
	logIt(msg, SS_LOG_LEVEL_WARN)
}

func LogError(msg string) {
	logIt(msg, SS_LOG_LEVEL_ERROR)
}

func LogFatal(msg string) {
	logIt(msg, SS_LOG_LEVEL_FATAL)
}

// Log our custom function entry event.
func LogEnter(msg ...string) {
	buf := DEBUG_TYPE_ENTER
	if len(msg) > 0 {
		buf = buf + " - " + strings.Join(msg, " ")
	}
	logIt(buf, SS_LOG_LEVEL_DEBUG)
}

// Log our custom function exit event.
func LogExit(msg ...string) {
	buf := DEBUG_TYPE_EXIT
	if len(msg) > 0 {
		buf = buf + " - " + strings.Join(msg, " ")
	}
	logIt(buf, SS_LOG_LEVEL_DEBUG)
}








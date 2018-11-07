// Utilities package.
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

// Logging configuration.
// NOTE: In production code report caller (Report_caller) is expensive and should be turned off.
// Provides two helper methods for logging function entry and exit.


var myLogFileHandle = initLogger()

func initLogger() (*os.File) {
	fmt.Println("simpleserver.util.logger - initLogger - ENTER")
	log.SetFlags(0)
	filename := MyConfig["log_file"]
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	} else {
		log.Fatal("Failed to open log file " + filename + ", using just stdout, ERROR: " + err.Error())
		log.SetOutput(os.Stdout)
	}
	fmt.Println("simpleserver.util.logger - initLogger - EXIT")
	return file
}

func CloseLog() {
	fmt.Println("simpleserver.util.logger - CloseLog - ENTER")
	myLogFileHandle.Close()
	fmt.Println("simpleserver.util.logger - CloseLog - EXIT")
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
var MyReportCaller = initReportCaller()


func initLogLevel() (SSLogLevel) {
	fmt.Println("simpleserver.util.logger - initLogLevel - ENTER")
	var myLogLevel SSLogLevel
	switch MyConfig["log_level"] {
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
		fmt.Println("simpleserver.util.logger.go - getLogLevel - ERROR: Unknown log level: " + MyConfig["log_level"])
		os.Exit(500)
	}
	fmt.Println("simpleserver.util.logger - initLogLevel - EXIT")
	return myLogLevel
}

func initReportCaller() (bool) {
	if MyConfig["report_caller"] == "true" {
		return true
	} else {
		return false
	}
}

// Provides string representation for log levels.
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
	var timeStamp = fmt.Sprint(time.Now().UTC().Format("2006-01-02T15:04:05.999Z"))
	if level >= MyLogLevel {
		if MyReportCaller {
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

// Log trace.
func LogTrace(msg string) {
	logIt(msg, SS_LOG_LEVEL_TRACE)
}

// Log debug.
func LogDebug(msg string) {
	logIt(msg, SS_LOG_LEVEL_DEBUG)
}

// Log info.
func LogInfo(msg string) {
	logIt(msg, SS_LOG_LEVEL_INFO)
}

// Log warning.
func LogWarn(msg string) {
	logIt(msg, SS_LOG_LEVEL_WARN)
}

// Log error.
func LogError(msg string) {
	logIt(msg, SS_LOG_LEVEL_ERROR)
}

// Log fatal.
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

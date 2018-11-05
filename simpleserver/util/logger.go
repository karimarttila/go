// Logging configuration package.
// NOTE: In production code report caller (Report_caller) is expensive and should be turned off.
// Provides three helper methods for logging function entry, exit and some arbitrary debug message.
// Application can use any logrus logging method (error, info...) using directly MyLogger public variable.

package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
)

// NOTE: In Go public variables and functions start with capital letter.
var MyLogger = initLogger()

func getLogLevel() (logrus.Level) {
	var myLogLevel logrus.Level
	switch MyConfig.Log_level {
	case "panic":
		myLogLevel = logrus.PanicLevel
	case "fatal":
		myLogLevel = logrus.FatalLevel
	case "error":
		myLogLevel = logrus.ErrorLevel
	case "warn":
		myLogLevel = logrus.WarnLevel
	case "info":
		myLogLevel = logrus.InfoLevel
	case "debug":
		myLogLevel = logrus.DebugLevel
	case "trace":
		myLogLevel = logrus.TraceLevel
	default:
		fmt.Println("simpleserver.util.logger.go - getLogLevel - ERROR: Unknown log level: " + MyConfig.Log_level)
		os.Exit(500)
	}
	return myLogLevel
}


func initLogger() (*logrus.Logger) {
	fmt.Println("simpleserver.util.logger.go - initLogger - ENTER")
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Formatter.(*logrus.TextFormatter).DisableColors = true
	log.Level = getLogLevel()
	// Let's handle this ourselves. We don't want the file name there.
	log.SetReportCaller(false)
	file, err := os.OpenFile(MyConfig.Log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	} else {
		log.Error("Failed to log to file, using default stderr")
	}
	fmt.Println("simpleserver.util.logger.go - initLogger - EXIT")
	return log
}


// Log our custom function entry event.
func LogEnter(msg ...string) {
	var myEntry *logrus.Entry
	if MyConfig.Report_caller {
		pc, _, _, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc)
		myEntry = MyLogger.WithFields(logrus.Fields{"debugtype": DEBUG_TYPE_ENTER, "caller": fn.Name()})
	} else {
		myEntry = MyLogger.WithFields(logrus.Fields{"debugtype": DEBUG_TYPE_ENTER})
	}
	myEntry.Debug(msg)
}


// Log our custom function exit event.
func LogExit(msg ...string) {
	var myEntry *logrus.Entry
	if MyConfig.Report_caller {
		// NOTE: Use underscore '_' when you don't need to reference certain return values.
		pc, _, _, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc)
		myEntry = MyLogger.WithFields(logrus.Fields{"debugtype": DEBUG_TYPE_EXIT, "caller": fn.Name()})
	} else {
		myEntry = MyLogger.WithFields(logrus.Fields{"debugtype": DEBUG_TYPE_EXIT})
	}
	myEntry.Debug(msg)
}


// Log some arbitrary function debug message.
func LogDebug(msg ...string) {
	var myEntry *logrus.Entry
	if MyConfig.Report_caller {
		pc, _, _, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc)
		myEntry = MyLogger.WithFields(logrus.Fields{"debugtype": DEBUG_TYPE_MSG, "caller": fn.Name()})
	} else {
		myEntry = MyLogger.WithFields(logrus.Fields{"debugtype": DEBUG_TYPE_MSG})
	}
	myEntry.Debug(msg)
}





package util

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var MyLogger = initLogger()

func initLogger() (*logrus.Logger) {
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Formatter.(*logrus.TextFormatter).DisableColors = true
	log.Level = logrus.TraceLevel
	// NOTE: In production code this is expensive and should be turned off.
	log.SetReportCaller(true)
	file, err := os.OpenFile("src/github.com/karimarttila/go/simpleserver/logs/simpleserver.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	return log
}

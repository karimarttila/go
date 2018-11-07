package util


import (
    "testing"
)


func TestLogger(t *testing.T) {
	LogEnter()
	if MyLogLevel != SS_LOG_LEVEL_TRACE {
		t.Error("Log level was not trace")
	}
	LogExit()
}

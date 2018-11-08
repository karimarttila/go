package util

import (
	"testing"
)

func TestMyConfig(t *testing.T) {
	LogEnter()
	if MyConfig["log_level"] != "trace" {
		t.Error("Error configuration not loaded correctly")
	}
	LogExit()
}

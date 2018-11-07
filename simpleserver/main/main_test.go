package main


import (
	"github.com/karimarttila/go/simpleserver/util"
	"testing"
)

// Dummy test.
func TestMainDummy(t *testing.T) {
	util.LogEnter()
	if false {
		t.Error("Error in main")
	}
	util.LogExit()
}

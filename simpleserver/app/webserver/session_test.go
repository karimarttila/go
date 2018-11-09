package webserver


import (
	"github.com/karimarttila/go/simpleserver/app/util"
	"testing"
)

func TestJsonWebToken(t *testing.T) {
	util.LogEnter()
	jsonWebToken := CreateJsonWebToken("kari.karttinen@foo.com")
	if len(jsonWebToken) < 20 {
		t.Errorf("jsonWebToken is too short")
	}
	util.LogExit()
}
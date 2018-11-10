package webserver


import (
	"github.com/karimarttila/go/simpleserver/app/util"
	"testing"
)

func TestJsonWebToken(t *testing.T) {
	util.LogEnter()
	testEmail := "kari.karttinen@foo.com"
	jsonWebToken, err := CreateJsonWebToken(testEmail)
	if err != nil {
		t.Error("CreateJsonWebToken returned error: " + err.Error())
	}
	if jsonWebToken == "" {
		t.Error("jsonWebToken is empty")
	}
	if len(jsonWebToken) < 20 {
		t.Error("jsonWebToken is too short")
	}
	response, err := ValidateJsonWebToken(jsonWebToken)
	if err != nil {
		t.Error("ValidateJsonWebToken returned error: " + err.Error())
	}
	if !response.Ready {
		t.Error("Couldn't validate token")
	}
	gotEmail := response.Email
	if gotEmail != testEmail {
		t.Errorf("Didn't get the same email back, expected: %s, got: %s", testEmail, gotEmail)
	}
	util.LogExit()
}


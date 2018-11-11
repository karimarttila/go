package userdb

import (
	"github.com/karimarttila/go/simpleserver/app/util"
	"strconv"
	"testing"
)

func TestEmailAlreadyExists(t *testing.T) {
	util.LogEnter()
	testingEmail := "kari.karttinen@foo.com"
	response := EmailAlreadyExists(testingEmail)
	if !response {
		t.Errorf("%s should have been found in the users DB", testingEmail)
	}
	testingEmail = "not.found@foo.com"
	response = EmailAlreadyExists(testingEmail)
	if response {
		t.Errorf("%s should not have been found in the users DB", testingEmail)
	}
	util.LogExit()
}

func TestAddUser(t *testing.T) {
	util.LogEnter()
	response, err := AddUser("kari.karttinen@foo.com", "Kari", "Karttinen", "Kari")
	if err == nil {
		t.Errorf("Adding user kari.karttinen@foo.com should have failed since it is in the user DB, response: %s", response)
	}
	response, err = AddUser("jamppa.jamppanen@foo.com", "Jamppa", "Jamppanen", "JampanSalasana")
	if err != nil {
		t.Errorf("Adding user jamppa.jamppanen@foo.com should have succeeded, response: %s", response)
	}
	if response.Ret != "ok" {
		t.Errorf("ret was something else than ok: %s", response.Ret)
	}
	util.LogExit()
}

func TestCheckCredentials(t *testing.T) {
	util.LogEnter()
	response := CheckCredentials("kari.karttinen@foo.com", "Kari")
	if !response {
		t.Errorf("User kari.karttinen@foo.com should have succeeded since both email and password ok, response: %s", strconv.FormatBool(response))
	}
	// Wrong password
	response = CheckCredentials("kari.karttinen@foo.com", "WRONG-PASSWORD")
	if response {
		t.Errorf("User kari.karttinen@foo.com should have failed since wrong password response: %s", strconv.FormatBool(response))
	}
	// Wrong email
	response = CheckCredentials("WRONG.USERNAME@foo.com", "Kari")
	if response {
		t.Errorf("User kari.karttinen@foo.com should have failed since wrong email, response: %s", strconv.FormatBool(response))
	}
	util.LogExit()
}

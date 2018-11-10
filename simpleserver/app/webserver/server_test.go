package webserver

import (
	"bytes"
	"encoding/json"
	"github.com/karimarttila/go/simpleserver/app/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetInfo(t *testing.T) {
	util.LogEnter()
	port := util.MyConfig["port"]
	//NOTE: We actually call directly the handler.
	// See below: "http.HandlerFunc(getInfo)...."
	request := httptest.NewRequest("GET", "http://localhost:"+port+"/info", nil)
	recorder := httptest.NewRecorder()
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("GetInfo handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(getInfo).ServeHTTP(recorder, request)
	response := recorder.Body.String()
	if len(response) == 0 {
		t.Error("Response was nil or empty")
	}
	retValue := strings.TrimRight(strings.TrimSpace(string(response[:])), "\n")
	util.LogDebug("/info returned: " + retValue)
	expectedValue := `{"info":"index.html => Info in HTML format"}`
	if retValue != expectedValue {
		t.Error("Expected value (" + expectedValue + ") != response value (" + retValue + ")")
	}
	util.LogEnter()
}

func addTestUser(t *testing.T) (recorder *httptest.ResponseRecorder, request *http.Request, testEmail string){
	util.LogEnter()
	port := util.MyConfig["port"]
	testEmail = "jamppa.jamppanen@foo.com"
    bodyMap := map[string]interface{}{
        "first-name": "Jamppa",
        "last-name": "Jamppanen",
        "email": testEmail,
        "password": "JampanSalasana",
    }
    myBody, _ := json.Marshal(bodyMap)
    request = httptest.NewRequest("POST", "http://localhost:"+port+"/signin", bytes.NewReader(myBody))
	recorder = httptest.NewRecorder()
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("PostSignin handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	util.LogEnter()
	return recorder, request, testEmail
}

func TestPostSignin(t *testing.T) {
	util.LogEnter()
	// First time adding the user, should go smoothly.
	recorder, request, testEmail := addTestUser(t)
	http.HandlerFunc(postSignin).ServeHTTP(recorder, request)
	responseStr := recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
    var responseMap map[string]interface{}
    err := json.NewDecoder(recorder.Body).Decode(&responseMap)
    if err != nil {
    	t.Errorf("Decoding request failed, err: %s", err.Error())
	}
    if responseMap["ret"] != "ok" {
    	t.Errorf("The response ret values should have been ok, map: %s", responseMap)
	}
    if responseMap["email"] != testEmail {
    	t.Errorf("The response didn't comprise the test email, map: %s", responseMap)
	}
	// Second time the user should be in the db already and signin should fail.
	recorder, request, testEmail = addTestUser(t)
	http.HandlerFunc(postSignin).ServeHTTP(recorder, request)
	responseStr = recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
    err = json.NewDecoder(recorder.Body).Decode(&responseMap)
    if err != nil {
    	t.Errorf("Decoding request failed, err: %s", err.Error())
	}
    if responseMap["ret"] != "failed" {
    	t.Errorf("The response ret values should have been failed, map: %s", responseMap)
	}
    if responseMap["msg"] != "Email already exists: " + testEmail {
    	t.Errorf("The response error msg was not correct, map: %s", responseMap)
	}
    if responseMap["email"] != testEmail {
    	t.Errorf("The response didn't comprise the test email, map: %s", responseMap)
	}
	util.LogEnter()
}



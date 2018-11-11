package webserver

import (
	"bytes"
	"encoding/base64"
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
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(getInfo).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getInfo handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
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

func addTestUser(t *testing.T, firstNameMissing bool) (recorder *httptest.ResponseRecorder, request *http.Request, testEmail string) {
	util.LogEnter()
	port := util.MyConfig["port"]
	testEmail = "jamppa.jamppanen@foo.com"
	bodyMap := map[string]interface{}{
		"last-name": "Jamppanen",
		"email":     testEmail,
		"password":  "JampanSalasana",
	}
	if !firstNameMissing {
		bodyMap["first-name"] = "Jamppa"
	}
	myBody, _ := json.Marshal(bodyMap)
	request = httptest.NewRequest("POST", "http://localhost:"+port+"/signin", bytes.NewReader(myBody))
	recorder = httptest.NewRecorder()
	util.LogEnter()
	return recorder, request, testEmail
}

func TestPostSignin(t *testing.T) {
	util.LogEnter()
	// Test missing parameter.
	var testEmail string
	recorder, request, _ := addTestUser(t, true)
	http.HandlerFunc(postSignin).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("postSignin handler returned wrong status code: expected: %v actual: %v",
			http.StatusBadRequest, status)
	}
	responseStr := recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
	var responseMap map[string]interface{}
	err := json.NewDecoder(recorder.Body).Decode(&responseMap)
	if err != nil {
		t.Errorf("Decoding request failed, err: %s", err.Error())
	}
	if responseMap["ret"] != "failed" {
		t.Errorf("The response ret value should have been 'failed', map: %s", responseMap)
	}
	if responseMap["msg"] != "Validation failed - some fields were empty" {
		t.Errorf("The validation should have comprised error message, map: %s", responseMap)
	}
	// First time adding the user, should go smoothly.
	recorder, request, testEmail = addTestUser(t, false)
	http.HandlerFunc(postSignin).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("postSignin handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	responseStr = recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
	err = json.NewDecoder(recorder.Body).Decode(&responseMap)
	if err != nil {
		t.Errorf("Decoding request failed, err: %s", err.Error())
	}
	if responseMap["ret"] != "ok" {
		t.Errorf("The response ret value should have been 'ok', map: %s", responseMap)
	}
	if responseMap["email"] != testEmail {
		t.Errorf("The response didn't comprise the test email, map: %s", responseMap)
	}
	// Second time the user should be in the db already and signin should fail.
	recorder, request, testEmail = addTestUser(t, false)
	http.HandlerFunc(postSignin).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("postSignin handler returned wrong status code: expected: %v actual: %v",
			http.StatusBadRequest, status)
	}
	responseStr = recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
	err = json.NewDecoder(recorder.Body).Decode(&responseMap)
	if err != nil {
		t.Errorf("Decoding request failed, err: %s", err.Error())
	}
	if responseMap["ret"] != "failed" {
		t.Errorf("The response ret value should have been 'failed', map: %s", responseMap)
	}
	if responseMap["msg"] != "Email already exists: "+testEmail {
		t.Errorf("The response error msg was not correct, map: %s", responseMap)
	}
	if responseMap["email"] != testEmail {
		t.Errorf("The response didn't comprise the test email, map: %s", responseMap)
	}
	util.LogEnter()
}

func TestLogin(t *testing.T) {
	util.LogEnter()
	port := util.MyConfig["port"]
	// First test failed login. Wrong password.
	bodyMap := map[string]interface{}{
		"email":    "kari.karttinen@foo.com",
		"password": "WRONG-PASSWORD",
	}
	myBody, _ := json.Marshal(bodyMap)
	request := httptest.NewRequest("POST", "http://localhost:"+port+"/login", bytes.NewReader(myBody))
	recorder := httptest.NewRecorder()
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(postLogin).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("postLogin handler returned wrong status code: expected: %v actual: %v",
			http.StatusBadRequest, status)
	}
	responseStr := recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
	var responseMap map[string]string
	err := json.NewDecoder(recorder.Body).Decode(&responseMap)
	if err != nil {
		t.Errorf("Decoding request failed, err: %s", err.Error())
	}
	if responseMap["ret"] != "failed" {
		t.Errorf("The response ret value should have been 'failed', map: %s", responseMap)
	}
	if responseMap["msg"] != "Credentials are not good - either email or password is not correct" {
		t.Errorf("The response msg was not correct, map: %s", responseMap)
	}
	// Then test ok login.
	//NOTE: We actually call directly the handler.
	// See below: "http.HandlerFunc(getInfo)...."
	bodyMap = map[string]interface{}{
		"email":    "kari.karttinen@foo.com",
		"password": "Kari",
	}
	myBody, _ = json.Marshal(bodyMap)
	request = httptest.NewRequest("POST", "http://localhost:"+port+"/login", bytes.NewReader(myBody))
	recorder = httptest.NewRecorder()
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(postLogin).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("postLogin handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	responseStr = recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
	err = json.NewDecoder(recorder.Body).Decode(&responseMap)
	if err != nil {
		t.Errorf("Decoding request failed, err: %s", err.Error())
	}
	if responseMap["ret"] != "ok" {
		t.Errorf("The response ret value should have been 'ok', map: %s", responseMap)
	}
	if responseMap["msg"] != "Credentials ok" {
		t.Errorf("The response msg was not correct, map: %s", responseMap)
	}
	jsonWebToken := responseMap["json-web-token"]
	if len(jsonWebToken) < 20 {
		t.Errorf("The json-web-token was too short, map: %s", responseMap)
	}
	util.LogEnter()
}

// Reusing the TestLogin functionality. Maybe refactoring later.
func getTestToken(t *testing.T) (token string, err error) {
	util.LogEnter()
	port := util.MyConfig["port"]
	bodyMap := map[string]interface{}{
		"email":    "kari.karttinen@foo.com",
		"password": "Kari",
	}
	myBody, _ := json.Marshal(bodyMap)
	request := httptest.NewRequest("POST", "http://localhost:"+port+"/login", bytes.NewReader(myBody))
	recorder := httptest.NewRecorder()
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(postLogin).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("postLogin handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	responseStr := recorder.Body.String()
	util.LogDebug("Got response: " + responseStr)
	var responseMap map[string]string
	err = json.NewDecoder(recorder.Body).Decode(&responseMap)
	if err != nil {
		t.Errorf("Decoding request failed, err: %s", err.Error())
	}
	if responseMap["ret"] != "ok" {
		t.Errorf("The response ret value should have been 'ok', map: %s", responseMap)
	}
	if responseMap["msg"] != "Credentials ok" {
		t.Errorf("The response msg was not correct, map: %s", responseMap)
	}
	token = responseMap["json-web-token"]
	if len(token) < 20 {
		t.Errorf("The json-web-token was too short, map: %s", responseMap)
	}
	util.LogEnter()
	return token, err
}

func TestGetProductGroups(t *testing.T) {
	util.LogEnter()
	port := util.MyConfig["port"]
	token, err := getTestToken(t)
	if err != nil {
		t.Errorf("Failed to get test token: %s", err.Error())
	}
	util.LogTrace("Test token: " + token)
	encoded := base64.StdEncoding.EncodeToString([]byte(token))
	if err != nil {
		t.Errorf("Failed to base64 decode token: %s", err.Error())
	}
	//NOTE: We actually call directly the handler.
	// See below: "http.HandlerFunc(getInfo)...."
	request := httptest.NewRequest("GET", "http://localhost:"+port+"/product-groups", nil)
	request.Header.Add("authorization", "Basic "+encoded)
	recorder := httptest.NewRecorder()
	// NOTE: Here we actually call directly the getInfo handler!
	http.HandlerFunc(getProductGroups).ServeHTTP(recorder, request)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getProductGroups handler returned wrong status code: expected: %v actual: %v",
			http.StatusOK, status)
	}
	response := recorder.Body.String()
	if len(response) == 0 {
		t.Error("Response was nil or empty")
	}
	// NOTE: Might look a bit weird, but it's pretty straightforward:
	// pgMap is a map (key:string), and values are maps, which keys are strings and values are strings.
	pgMap := make(map[string]map[string]string)
	err = json.Unmarshal([]byte(response), &pgMap)
	if err != nil {
		t.Errorf("Unmarshalling response failed: %s", err.Error())
	}
	pg, ok := pgMap["product-groups"]
	if !ok {
		t.Errorf("Didn't find 'product-groups' in response")
	}
	pg1, ok := pg["1"]
	if !ok {
		t.Errorf("Didn't find product group 1 in response")
	}
	if pg1 != "Books" {
		t.Errorf("Product group 1 should have been 'Books'")
	}
	util.LogEnter()
}

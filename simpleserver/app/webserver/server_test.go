package webserver

import (
	"github.com/karimarttila/go/simpleserver/app/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetInfo(t *testing.T) {
	util.LogEnter()
	port := util.MyConfig["port"]
	request := httptest.NewRequest("GET", "http://localhost:"+port+"info", nil)
	recorder := httptest.NewRecorder()
	if status := recorder.Code; status != http.StatusOK {
        t.Errorf("GetInfo handler returned wrong status code: expected: %v actual: %v",
            http.StatusOK, status)
    }
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

//{"product-groups":{"1":"Books","2":"Movies"},"ret":"ok"}

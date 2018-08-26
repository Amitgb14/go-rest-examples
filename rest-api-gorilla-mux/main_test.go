package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkReponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user/2", nil)
	response := executeRequest(req)

	checkReponseCode(t, http.StatusOK, response.Code)
	// var m map[string]interface{}
	// json.Unmarshal(response.Body.Bytes(), &m)
}

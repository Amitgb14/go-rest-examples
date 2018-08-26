package main

import (
	"bytes"
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

func TestAddUser(t *testing.T) {
	payload := []byte(`{"ID": 3, "Name": "Ankit", "Age": 21}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkReponseCode(t, http.StatusCreated, response.Code)

}

func TestAddFile(t *testing.T) {

	data := []byte(`
	<Data>
		<ID>4</ID>
		<Name>Grace R. Emlin</Name>
		<Age>35</Age>
	</Data>`)
	req, _ := http.NewRequest("POST", "/file", bytes.NewBuffer(data))
	response := executeRequest(req)

	checkReponseCode(t, http.StatusCreated, response.Code)

}

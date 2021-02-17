package tests

import (
	"api"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var a api.App

func TestMain(m *testing.M) {
	a = api.App{}
	a.Initialise()
	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestInsertRetrieve(t *testing.T) {
	r := strings.NewReader("my request")
	req, _ := http.NewRequest("POST", "/queue/pau", r)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/queue/pau", nil)
	response = executeRequest(req)

	if body := response.Body.String(); body != "my request" {
		t.Errorf("Expected my request. Got %s", body)
	}
}

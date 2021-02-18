package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
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

func TestInsertOneInt(t *testing.T) {
	r := strings.NewReader("my request")
	req, _ := http.NewRequest("POST", "/queue/pau", r)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	length := a.QueuesMananger.Len("pau")
	if length != 1 {
		t.Errorf("Expected queue leng is 1 and not %d", length)
	}

	req, _ = http.NewRequest("GET", "/queue/pau", nil)
	response = executeRequest(req)

	if body := response.Body.String(); body != "my request" {
		t.Errorf("Expected my request. Got %s", body)
	}
}

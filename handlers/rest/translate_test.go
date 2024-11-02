package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tuminzee/hello-api/handlers/rest"
)

func TestTranslate(t *testing.T) {

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	handler := http.HandlerFunc(rest.TranslateHandler)

	handler.ServeHTTP(rr, req)

	t.Logf("Response body: %s", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	var resp rest.Resp
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.Language != "english" {
		t.Errorf("handler returned wrong language: got %v want %v", resp.Language, "English")
	}

	if resp.Translation != "hello" {
		t.Errorf("handler returned wrong translation: got %v want %v", resp.Translation, "hello")
	}
}

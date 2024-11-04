package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tuminzee/hello-api/handlers/rest"
	"github.com/tuminzee/hello-api/translation"
)

func TestGetLanguages(t *testing.T) {
	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/languages", nil)
	handler := http.HandlerFunc(rest.LanguagesHandler)

	handler.ServeHTTP(rr, req)

	t.Logf("Response body: %s", rr.Body.Bytes())

	var resp rest.LanguagesResp
	json.Unmarshal(rr.Body.Bytes(), &resp)

	languages := translation.GetSupportedLanguages()

	if len(resp.Languages) != len(languages) {
		t.Errorf("Expected languages: %v, got: %v", resp.Languages, rr.Body.String())
	}

	t.Logf("Response languages: %v", resp.Languages)
}

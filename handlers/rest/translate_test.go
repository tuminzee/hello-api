package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tuminzee/hello-api/handlers/rest"
)

func TestTranslate(t *testing.T) {

	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=french",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "french",
			ExpectedTranslation: "bonjour",
		},
		{
			Endpoint:            "/hello?language=spanish",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "spanish",
			ExpectedTranslation: "hola",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=unknown",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	for _, tc := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", tc.Endpoint, nil)
		handler := http.HandlerFunc(rest.TranslateHandler)

		handler.ServeHTTP(rr, req)

		t.Logf("Response body: %s", rr.Body.String())

		if rr.Code != tc.StatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tc.StatusCode)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		t.Logf("Responce Word: %s, Response language: %s", resp.Translation, resp.Language)

		if resp.Language != tc.ExpectedLanguage {
			t.Errorf("handler returned wrong language: got %v want %v", resp.Language, tc.ExpectedLanguage)
		}

		if resp.Translation != tc.ExpectedTranslation {
			t.Errorf("handler returned wrong translation: got %v want %v", resp.Translation, tc.ExpectedTranslation)
		}
	}
}

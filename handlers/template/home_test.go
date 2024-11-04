package template

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHome_TemplateNotFound(t *testing.T) {
	// Temporarily rename or move the template file to simulate missing template
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleHome)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check if we get an internal server error when template is not found
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

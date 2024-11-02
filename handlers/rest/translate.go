package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/tuminzee/hello-api/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := defaultLanguage

	word := strings.TrimPrefix(r.URL.Path, "/")
	slog.Info("got the word", "word", word)

	translation := translation.Translate(word, language)

	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package rest

import (
	"encoding/json"
	"net/http"

	"github.com/tuminzee/hello-api/translation"
)

type LanguagesResp struct {
	Languages []string `json:"languages"`
}

func LanguagesHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	languages := translation.GetSupportedLanguages()
	err := enc.Encode(LanguagesResp{Languages: languages})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

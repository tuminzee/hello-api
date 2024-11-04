package main

import (
	"log"

	"net/http"

	"github.com/tuminzee/hello-api/handlers/rest"
	"github.com/tuminzee/hello-api/handlers/template"
)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/", template.HandleHome)
	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/languages", rest.LanguagesHandler)
	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

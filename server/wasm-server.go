package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("wasm"))
	http.Handle("/", fs)

	log.Print("Serving WASM at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

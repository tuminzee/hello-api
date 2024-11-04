package template

import (
	"net/http"
	"text/template"
)

type HomeData struct {
	Title string
}

func HandleHome(w http.ResponseWriter, _ *http.Request) {
	homeData := HomeData{
		Title: "Hello API",
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, homeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package server

import (
	"html/template"
	"net/http"

	asciiart "asciiartweb/internal/ascii-art"
)

type output struct {
	Ascii_art, Text string
}

func handler_ascii_art(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		returnError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	fs := r.FormValue("artOption")
	if err := asciiart.AsciiChar(text); err != nil {
		returnError(w, http.StatusBadRequest)
		return
	} else {
		ascii_art, err := asciiart.Build(text, fs)
		if err != nil {
			returnError(w, http.StatusInternalServerError)
			return
		}
		out := output{Ascii_art: ascii_art, Text: text}
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			returnError(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, out)
		if err != nil {
			returnError(w, http.StatusInternalServerError)
			return
		}
	}
}

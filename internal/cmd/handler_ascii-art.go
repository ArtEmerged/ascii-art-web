package server

import (
	"net/http"

	asciiart "asciiartweb/internal/ascii-art"
)

type output struct {
	Ascii_art, Text string
}

func handler_ascii_art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("data")
	banners := r.FormValue("banners")
	ascii_art, code := asciiart.Build(text, banners)
	if code != 0 {
		returnError(w, code)
		return
	}
	out := output{
		Ascii_art: ascii_art,
		Text:      text}
	err := Tmpl.ExecuteTemplate(w, "index.html", out)
	if err != nil {
		returnError(w, http.StatusInternalServerError)
		return
	}

}

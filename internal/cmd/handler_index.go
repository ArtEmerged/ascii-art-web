package server

import (
	"net/http"
)

func handler_index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		returnError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
	err := Tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		returnError(w, http.StatusInternalServerError)
		return
	}
}

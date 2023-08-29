package server

import (
	"html/template"
	"net/http"
)

func handler_index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		returnError(w, http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			returnError(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			returnError(w, http.StatusInternalServerError)
			return
		}
	default:
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
}

package server

import (
	"fmt"
	"net/http"
)

type Err struct {
	Text_err string
	Code_err int
}

func returnError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	res := &Err{Text_err: http.StatusText(code), Code_err: code}
	err := Tmpl.ExecuteTemplate(w, "error.html", *res)
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
}

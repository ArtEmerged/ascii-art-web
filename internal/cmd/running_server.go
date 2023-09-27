package server

import (
	"html/template"
	"log"
	"net/http"
)

var Tmpl *template.Template

func init() {
	Tmpl = template.Must(template.ParseGlob("web/templates/*.html"))
}

func Running_server() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	http.HandleFunc("/", handler_index)
	http.HandleFunc("/ascii-art", handler_ascii_art)
	log.Println("Listening on: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

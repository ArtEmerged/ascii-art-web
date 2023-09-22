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
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	mux.HandleFunc("/", handler_index)
	mux.HandleFunc("/ascii-art", handler_ascii_art)
	log.Println("Listening on: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

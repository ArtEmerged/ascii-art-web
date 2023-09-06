package server

import (
	"log"
	"net/http"
)

func Running_server() {
	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))
	mux.HandleFunc("/", handler_index)
	mux.HandleFunc("/ascii-art", handler_ascii_art)
	log.Println("Listening on: http://localhost:5080/")
	log.Fatal(http.ListenAndServe(":5080", mux))
	// httpServer := exec.Command("xdg-open", "http://localhost:8080")
	// if err := httpServer.Run(); err != nil {
	// 	panic(err)
	// }
}

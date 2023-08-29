package server

import (
	"log"
	"net/http"
)

func Running_server() {
	// init CSS
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler_index)
	mux.HandleFunc("/ascii-art", handler_ascii_art)
	log.Println("Listening on: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
	// httpServer := exec.Command("xdg-open", "http://localhost:8080")
	// if err := httpServer.Run(); err != nil {
	// 	panic(err)
	// }
}

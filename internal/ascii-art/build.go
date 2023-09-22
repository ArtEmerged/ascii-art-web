package asciiart

import (
	"log"
	"net/http"
)

func Build(text, fs string) (string, int) {
	if err := AsciiChar(text); err != nil {
		log.Println(err)
		return "", http.StatusBadRequest
	}
	if fs != "standard" && fs != "shadow" && fs != "thinkertoy" {
		log.Println("User requests an incorrect banner")
		return "", http.StatusBadRequest
	}
	if err := Validate(); err != nil {
		log.Println(err)
		return "", http.StatusInternalServerError
	}
	fsChar, err := AsciiList(fs)
	if err != nil {
		log.Println(err)
		return "", http.StatusInternalServerError
	}

	split_text := Split(text)
	return Print(split_text, fsChar), 0
}

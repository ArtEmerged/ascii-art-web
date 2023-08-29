package asciiart

import "errors"

const (
	firstCharAscii = rune(32)
	lastCharAscii = rune(126)
)

func AsciiChar(text string) error {
	if text == "" {
		return errors.New("empty response")
	}
	for _, char := range text {
		if char == rune(13) {
			char = rune(10)
		}
		if !(char >= firstCharAscii && char <= lastCharAscii || char == rune(10)) {
			return errors.New("an invalid character was received in the response")
		}
	}
	return nil
}

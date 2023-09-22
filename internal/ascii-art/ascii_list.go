package asciiart

import (
	"bufio"
	"os"
)

const (
	firstRuneMin1 = rune(31)
)

func AsciiList(file string) (map[rune][]string, error) {
	letter := make(map[rune][]string)
	in, err := os.Open("internal/ascii-art/banners/" + file + ".txt")
	f := bufio.NewScanner(in)
	char := firstRuneMin1
	for f.Scan() {
		if f.Text() == "" {
			char++
			continue
		}
		letter[char] = append(letter[char], f.Text())
	}
	return letter, err
}

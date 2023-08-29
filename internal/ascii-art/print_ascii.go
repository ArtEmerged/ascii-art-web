package asciiart

const (
	lineInText_0 = 0
	lineInText_8 = 8
	lastLine     = 7
)

func Print(split_text []string, fsChar map[rune][]string) string {
	result := ""
	for _, word := range split_text {
		if word == "\n" {
			result += "\n"
			continue
		} else {
			for lineInText := lineInText_0; lineInText < lineInText_8; lineInText++ {
				for _, char := range word {
					result += fsChar[char][lineInText]
				}
				if lineInText != lastLine {
					result += "\n"
				}
			}
		}
	}
	return result
}

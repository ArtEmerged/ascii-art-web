package asciiart

func Build(text, fs string) (string, error) {
	if err := AsciiChar(text); err != nil {
		return "", err
	}
	if err := Validate(); err != nil {
		return "", err
	}
	fsChar, err := AsciiList(fs)
	if err != nil {
		return "", err
	}

	split_text := Split(text)
	return Print(split_text, fsChar), nil
}

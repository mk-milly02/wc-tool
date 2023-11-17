package ccwc

import (
	"bytes"
	"strings"
)

func GetNumberOfBytes(text []byte) int {
	return len(text)
}

func GetNumberOfLines(text []byte) int {
	reader := bytes.NewBuffer(text)
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return len(lines)
}

func GetNumberOfWords(text []byte) int {
	words := strings.Fields(string(text))
	return len(words)
}

func GetNumberOfCharacters(text []byte) int {
	return len([]rune(string(text)))
}

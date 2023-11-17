package ccwc

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func GetNumberOfBytes(filepath string) int {
	file, closer := openFile(filepath)
	defer closer()

	filebytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len(filebytes)
}

func GetNumberOfLines(filepath string) int {
	file, closer := openFile(filepath)
	defer closer()

	reader := bufio.NewReader(file)
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

func GetNumberOfWords(filepath string) int {
	file, closer := openFile(filepath)
	defer closer()

	filebytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(filebytes))
	return len(words)
}

func GetNumberOfCharacters(filepath string) int {
	file, closer := openFile(filepath)
	defer closer()

	filebytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len([]rune(string(filebytes)))
}

func openFile(filepath string) (*os.File, func()) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("%s: no such file or directory", filepath)
	}
	return file, func() { file.Close() }
}

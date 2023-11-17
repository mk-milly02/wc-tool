package ccwc

import (
	"io"
	"log"
	"os"
)

func GetNumberOfCharacters(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("%s: no such file or directory", filepath)
	}
	defer file.Close()

	filebytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len(filebytes)
}

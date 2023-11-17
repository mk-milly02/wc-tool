package ccwc_test

import (
	"io"
	"log"
	"os"
	"testing"
	"wctool/ccwc"
)

func TestGetNumberOfBytes(t *testing.T) {
	t.Parallel()
	want := 342190
	got := ccwc.GetNumberOfBytes(openFile("../test.txt"))
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetNumberOfNewLines(t *testing.T) {
	t.Parallel()
	want := 7145
	got := ccwc.GetNumberOfLines(openFile("../test.txt"))
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetNumberOfWords(t *testing.T) {
	t.Parallel()
	want := 58164
	got := ccwc.GetNumberOfWords(openFile("../test.txt"))
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetNumberOfCharacters(t *testing.T) {
	t.Parallel()
	want := 339292
	got := ccwc.GetNumberOfCharacters(openFile("../test.txt"))
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func openFile(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("%s: no such file or directory", filepath)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

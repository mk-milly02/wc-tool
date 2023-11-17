package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"wctool/ccwc"
)

func main() {
	c := flag.Bool("c", false, "print the byte counts")
	l := flag.Bool("l", false, "print the newline counts")
	w := flag.Bool("w", false, "print the word counts")
	m := flag.Bool("m", false, "print the character counts")

	flag.Parse()

	var filename string
	var text []byte

	if filename = flag.Arg(0); filename != "" {
		text = readFromFile(filename)
	} else {
		text = readFromStdIn()
	}

	if *c {
		numOfBytes := ccwc.GetNumberOfBytes(text)
		fmt.Printf("  %v %s", numOfBytes, filename)
	} else if *l {
		numOfLines := ccwc.GetNumberOfLines(text)
		fmt.Printf("  %v %s", numOfLines, filename)
	} else if *w {
		numOfWords := ccwc.GetNumberOfWords(text)
		fmt.Printf("  %v %s", numOfWords, filename)
	} else if *m {
		numOfChars := ccwc.GetNumberOfCharacters(text)
		fmt.Printf("  %v %s", numOfChars, filename)
	} else {
		numOfLines := ccwc.GetNumberOfLines(text)
		numOfWords := ccwc.GetNumberOfWords(text)
		numOfBytes := ccwc.GetNumberOfBytes(text)
		fmt.Printf("  %v %v %v %s", numOfLines, numOfWords, numOfBytes, filename)
	}
}

func readFromFile(filepath string) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("  %s: no such file or directory", filepath)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func readFromStdIn() []byte {
	reader := bufio.NewReader(os.Stdin)
	content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return []byte(content)
}

package main

import (
	"flag"
	"fmt"
	"wctool/ccwc"
)

func main() {
	c := flag.Bool("c", false, "print the byte counts")
	l := flag.Bool("l", false, "print the newline counts")
	w := flag.Bool("w", false, "print the word counts")
	m := flag.Bool("m", false, "print the character counts")

	filepath := flag.String("file", "test.txt", "")

	flag.Parse()

	if *c {
		numOfBytes := ccwc.GetNumberOfBytes(*filepath)
		fmt.Printf("%v %s", numOfBytes, *filepath)
	}

	if *l {
		numOfLines := ccwc.GetNumberOfLines(*filepath)
		fmt.Printf("%v %s", numOfLines, *filepath)
	}

	if *w {
		numOfWords := ccwc.GetNumberOfWords(*filepath)
		fmt.Printf("%v %s", numOfWords, *filepath)
	}

	if *m {
		numOfChars := ccwc.GetNumberOfCharacters(*filepath)
		fmt.Printf("%v %s", numOfChars, *filepath)
	}
}

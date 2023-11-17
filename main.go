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

	filepath := flag.String("file", "test.txt", "")

	flag.Parse()

	if *c {
		numOfChars := ccwc.GetNumberOfCharacters(*filepath)
		fmt.Printf("%v %s", numOfChars, *filepath)
	}

	if *l {
		numOfLines := ccwc.GetNumberOfLines(*filepath)
		fmt.Printf("%v %s", numOfLines, *filepath)
	}

	if *w {
		numOfWords := ccwc.GetNumberOfWords(*filepath)
		fmt.Printf("%v %s", numOfWords, *filepath)
	}
}

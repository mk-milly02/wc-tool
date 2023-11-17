package main

import (
	"flag"
	"fmt"
	"wc/ccwc"
)

func main() {
	m := flag.Bool("m", true, "print the byte counts")
	filepath := flag.String("file", "test.txt", "")
	flag.Parse()

	if *m {
		filebytes := ccwc.GetNumberOfCharacters(*filepath)
		fmt.Printf("%v %s", filebytes, *filepath)
	}
}

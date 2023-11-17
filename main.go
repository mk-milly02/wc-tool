package main

import (
	"flag"
	"fmt"
	"wc/ccwc"
)

func main() {
	c := flag.Bool("c", true, "print the byte counts")
	filepath := flag.String("file", "test.txt", "")
	flag.Parse()

	if *c {
		filebytes := ccwc.GetNumberOfCharacters(*filepath)
		fmt.Printf("%v %s", filebytes, *filepath)
	}
}

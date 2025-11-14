package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// TODO: how to test the main (specially the flags)?
func main() {
	var output int

	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()

	if *lines {
		output = countLines(os.Stdin)
	} else {
		output = count(os.Stdin)
	}

	fmt.Println(output)
}

// TODO: understand:
// - io.Reader
// - bufio.NewScanner
func count(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}

func countLines(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}

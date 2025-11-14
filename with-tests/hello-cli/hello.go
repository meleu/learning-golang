package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(Hello(name))
}

func Hello(name string) string {
	var greeting string
	lang := os.Getenv("LANG")

	if name == "" {
		name = "World"
	}

	if strings.HasPrefix(lang, "fr_") {
		greeting = "Bonjour, "
	} else {
		greeting = "Hello, "
	}

	return greeting + name
}

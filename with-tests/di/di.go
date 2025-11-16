package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Greet(os.Stdout, "Elodie")
	log.Fatal(
		http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)),
	)
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

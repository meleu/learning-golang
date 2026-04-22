package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "meleu")

	got := buffer.String()
	want := "Hello, meleu"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

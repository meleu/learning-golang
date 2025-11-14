package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("word1 word2 word3 \nword4\nword5 word6")
	want := 6
	got := count(input)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("word1 word2 word3 \nword4\nword5 word6")
	want := 3
	got := countLines(input)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

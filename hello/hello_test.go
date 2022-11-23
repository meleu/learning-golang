package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("meleu")
  want := "Hello, meleu"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("meleu")
	expected := "Hello, meleu"

	if actual != expected {
		t.Errorf("actual '%s', expected '%s'", actual, expected)
	}
}

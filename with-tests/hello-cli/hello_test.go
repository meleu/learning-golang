package main

import (
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("in French when 'LANG=fr_*'", func(t *testing.T) {
		originalLang := os.Getenv("LANG")

		os.Setenv("LANG", "fr_FR.UTF-8")
		actual := Hello("Jean")
		expected := "Bonjour, Jean"
		assertEqualString(t, actual, expected)

		os.Setenv("LANG", originalLang)
	})

	t.Run("say hello to people", func(t *testing.T) {
		actual := Hello("John")
		expected := "Hello, John"
		assertEqualString(t, actual, expected)
	})

	t.Run("say 'Hello, World' when no name is given", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World"
		assertEqualString(t, actual, expected)
	})
}

func assertEqualString(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("\n  actual: %q\nexpected: %q", actual, expected)
	}
}

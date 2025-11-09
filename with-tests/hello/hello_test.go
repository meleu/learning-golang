package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("John", "")
		expected := "Hello, John"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Juan", "Spanish")
		expected := "Hola, Juan"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in French", func(t *testing.T) {
		actual := Hello("Jean", "French")
		expected := "Bonjour, Jean"
		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("\n  actual: %q\nexpected: %q", actual, expected)
	}
}

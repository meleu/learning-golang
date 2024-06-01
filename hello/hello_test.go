package main

// needed to use `*testing.T`
import "testing"

// About the test function
// - must start with `Test`
// - takes only one argument `t *testing.T`
// - `t` is your "hook" into the testing framework
func TestHello(t *testing.T) {
	// 👇 t.Run(testName, testFunction)
	t.Run("saying 'Olá' to people", func(t *testing.T) {
		got := Hello("meleu", "Portuguese")
		want := "Olá, meleu"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Bonjour' to people", func(t *testing.T) {
		got := Hello("meleu", "French")
		want := "Bonjour, meleu"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hola' to people", func(t *testing.T) {
		got := Hello("meleu", "Spanish")
		want := "Hola, meleu"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("meleu", "")
		want := "Hello, meleu"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

/*
For helper functions, it's a good idea to accept a testing.TB
which is an interface that *testing.T and *testing.B both satisfy,
so you can call helper functions from a test, or a benchmark.
*/
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper(): needed to tell the test suite assertCorrectMessage is a helper.
	// This way, when it fails the line number reported will be in the caller.
	t.Helper()
	if got != want {
		// `t.Errorf` prints a message when a test fails.
		t.Errorf("got %q want %q", got, want)
	}
}

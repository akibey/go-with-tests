package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Aniruddha")
		want := "Hello, Aniruddha"

        assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults ot 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

        assertCorrectMessage(t, got, want)
	})
}

// Avoid duplication of the if statement

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

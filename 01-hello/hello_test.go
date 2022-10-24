package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Aniruddha")
	want := "Hello, Aniruddha"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

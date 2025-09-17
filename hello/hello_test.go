package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to specific people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		AssertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world when no string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sylvie", "French")
		want := "Bonjour, Sylvie"
		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

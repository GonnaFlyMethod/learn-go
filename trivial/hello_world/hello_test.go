package hello_world

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	t.Run("saying hello_world to people", func(t *testing.T) {
		got := hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("no name of a man, hello_world world expected", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello_world using spanish", func(t *testing.T) {
		got := hello("John", "Spanish")
		want := "Hola, John"
		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello_world using French", func(t *testing.T) {
		got := hello("John", "French")
		want := "Bonjure, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello_world using Russian", func(t *testing.T) {
		got := hello("John", "Russian")
		want := "Privet, John"
		assertCorrectMessage(t, got, want)
	})

}

package main

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()

		if got != want{
			t.Errorf("Got %q want %q",got, want)
		}
	}

	t.Run("simple adding" , func(t *testing.T){
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

}

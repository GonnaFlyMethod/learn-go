package integers

import "testing"


func TestAdder(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got, want int){
		t.Helper()

		if got != want{
			t.Errorf("Got %q want %q",got, want)
		}
	}

	t.Run("simple adding" , func(t *testing.T){
		got := Add(4, 4)
		want := 8

		assertCorrectMessage(t, got, want)
	})

}
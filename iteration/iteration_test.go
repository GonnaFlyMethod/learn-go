package itteration

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

func BenchmarkRepeat(b *testing.B) {
	var N int = 1000000
	for i := 0 ;i < N; i++{
		Repeat("a", 50)
	}
}

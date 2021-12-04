package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		got := Sum(numbers)
		want := 36

		if got != want {
			t.Errorf("Got: %d, want: %d, values: %v", got, want, numbers)
		}
	})

}
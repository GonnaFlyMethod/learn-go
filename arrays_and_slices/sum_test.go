package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {

	t.Run("Basic case", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2, 0})
		want := []int{6, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	}

	t.Run("basic case", func(t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{5, 4, 2, 8})
		want := []int{2, 14}

		checkSums(t, got, want)
	})

	t.Run("safely handling of empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 2}, []int{1, 2, 3, 4})
		want := []int{0, 2, 9}

		checkSums(t, got, want)
	})
}

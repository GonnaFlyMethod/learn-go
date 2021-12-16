package slicing_in_go

import (
	"reflect"
	"testing"
)

func TestChangeAllElementsThatGetUnderSlice(t *testing.T) {
	t.Run("Common case", func(t *testing.T) {
		got := ChangeAllElementsThatGetUnderSlice([]int{1, 2, 3, 4, 5, 6, 7}, 3, 5, 1000)
		want := []int{1, 2, 3, 1000, 1000, 6, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})
}

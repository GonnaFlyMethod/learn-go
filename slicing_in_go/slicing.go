package slicing_in_go

import "github.com/go-chi/chi/v5"

func ChangeAllElementsThatGetUnderSlice(sequence []int, sliceStart, sliceEnd, value int) []int {
	var sequenceBeginning, partToChange, sequenceEnd []int

	sequenceBeginning = sequence[:sliceStart]
	partToChange = sequence[sliceStart:sliceEnd]
	sequenceEnd = sequence[sliceEnd:]
	chi.NewMux()
	for i := 0; i < len(partToChange); i++ {
		partToChange[i] = value
	}

	semiResult := append(sequenceBeginning, partToChange...)
	return append(semiResult, sequenceEnd...)
}

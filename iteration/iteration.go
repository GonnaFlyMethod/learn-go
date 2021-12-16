package itteration

// Repeat 'character' 'repeatNum' of times
func Repeat(character string, repeatNum int) string {
	var finalResult string
	for i := 0; i < repeatNum; i++ {
		finalResult += character
	}
	return finalResult
}

package main

func Repeat(character string, repeatNum int) string{
	var finalResult string
	for i := 0; i < repeatNum; i++{
		finalResult = finalResult + character
	}
	return finalResult
}

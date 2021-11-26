package main

import "fmt"

const helloPrefixEnglishEdition = "Hello, "
const helloPrefixFrenchEdition = "Bonjure, "
const helloPrefixSpanishEdition = "Hola, "
const helloPrefixRussianEdition = "Privet, "


func choosePrefixForLang(lang string) string{
	switch lang {
	case "Spanish":
		return helloPrefixSpanishEdition
	case "French":
		return helloPrefixFrenchEdition
	case "Russian":
		return helloPrefixRussianEdition
	default:
		return helloPrefixEnglishEdition
	}
}


func hello(name , language string) (finalRes string) {
	prefix := choosePrefixForLang(language)

	if name == ""{
		finalRes = prefix + "World"
		return
	}

	finalRes = prefix + name
	return
}


func main(){
	stringToPrint := hello("John", "Russian")
	fmt.Println(stringToPrint)
}

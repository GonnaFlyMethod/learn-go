package hello_world

const helloPrefixRussianEdition = "Privet, "
const helloPrefixEnglishEdition = "Hello, "
const helloPrefixFrenchEdition = "Bonjure, "
const helloPrefixSpanishEdition = "Hola, "

func choosePrefixForLang(lang string) string {
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

func hello(name, language string) (finalRes string) {
	prefix := choosePrefixForLang(language)

	if name == "" {
		finalRes = prefix + "World"
		return
	}

	finalRes = prefix + name
	return
}

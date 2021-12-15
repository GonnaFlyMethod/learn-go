package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not found the word you were looking for")
	ErrWordExists       = DictionaryErr("can't add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

package maps

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Dictionary custom type", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Dictionary unknown word", func(t *testing.T) {
		_, err := dictionary.Search("Unknown")

		if err == nil {
			t.Fatal("Wanted err but get nil")
		}

		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("such value doesn't exist yet", func(t *testing.T) {
		key, value := "test", "this is just a test"

		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("such value already exist", func(t *testing.T) {
		key, value := "test", "this is just a test"

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("A word exists in dict", func(t *testing.T) {
		word, definition := "test", "this is just a test"

		dictionary := Dictionary{word: definition}

		newDefinition := "new definition"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("A word doesn't exists in dict", func(t *testing.T) {
		word, definition := "test", "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update("test1", definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "this is just a test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertError(t testing.TB, gotError error, wantedError error) {
	if gotError != wantedError {
		t.Fatalf("got %q, want %q", gotError, wantedError)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != definition {
		t.Errorf("got %q, want %q", got, definition)
	}

}

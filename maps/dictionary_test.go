package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a test",
	}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("something else")
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "word"
		description := "this is an other description"
		err := dictionary.Add(word, description)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, description)
	})
	t.Run("existing word", func(t *testing.T) {

		word := "word"
		description := "this is an other description"
		dictionary := Dictionary{
			word: description,
		}
		err := dictionary.Add(word, "new description")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, description)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, description string) {
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, description)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound   = DictionaryErr("could not find that word")
	ErrWordExists = DictionaryErr("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, description string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = description
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

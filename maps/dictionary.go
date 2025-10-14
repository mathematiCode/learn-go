package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot perform operation on word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	} else {
		return definition, nil
	}
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	} else {
		d[word] = definition
		return nil
	}
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, ok := d[word]
	if ok {
		d[word] = newDefinition
		return nil
	} else {
		return ErrWordDoesNotExist
	}
}

func (d Dictionary) Delete(word string) error {
	_, ok := d[word]
	if ok {
		delete(d, word)
		return nil
	} else {
		return ErrWordDoesNotExist
	}
}

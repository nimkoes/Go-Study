package myDict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {

	// [if style]
	_, err := d.Search(word)

	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}

	return nil

	// [switch style]
	/*
		_, err := d.Search(word)

		switch err {
		case errNotFound:
			d[word] = def
		case nil:
			return errWordExists
		}

		return nil
	*/
}

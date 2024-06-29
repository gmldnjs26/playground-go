package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("not found")
	errWordExists = errors.New("already exists")
	errNotExists  = errors.New("not exists")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errNotExists
	case nil:
		d[word] = def
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errNotExists
	case nil:
		delete(d, word)
	}

	return nil
}

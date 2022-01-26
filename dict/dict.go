package dict

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
	errExists = errors.New("Exists")
	errCantUpdate = errors.New("Cannot Update Non-Existing Word")
	errCantDelete = errors.New("Cannot Delete Non-Existing Wrod")
) 

func (d Dictionary) Search(target string) (string, error) {
	value, ok := d[target]
	if ok {
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
		return errExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
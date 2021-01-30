package maps

//Dictionary is a new dictionary
type Dictionary map[string]string

//DictionaryErr used to return an error
type DictionaryErr string

//ErrNotFound value to use when dictionary term not found
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists in dict")
	ErrWordDoesNotExist = DictionaryErr("word to update doesnt exist in dict")
)

//Search this is a search function
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add allows user to add a term to our Map
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

//Update allows user to update an existing dictionary definition
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

//Delete allows user to delete an existing definition
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}

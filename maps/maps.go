package maps

//Dictionary is a new dictionary
type Dictionary map[string]string

//Search this is a search function
func (d Dictionary) Search(word string) string {
	return d[word]
}

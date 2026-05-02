package maps
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)
type DictionaryErr string
func ( e DictionaryErr) Error()string {
	return string(e)
}
type Dictionary map[string] string
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil

}
func (d Dictionary) Add(key string, value string)error {
	_,err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key]=value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
	
	
}

// func Search(dictionary map[string]string, word string )string {
// 	return dictionary[word]
// }

func (d Dictionary) Update(key, value string)error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}


func (d Dictionary) Delete(key string)error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil


}
package myDict

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("That word already exists")
)

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

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		// 단어를 찾는데 err 가 nil -> 찾는데 성공했으므로 그 값을 업데이트
		d[word] = definition
	case errNotFound:
		// 단어를 찾는데 실패 -> 존재하지 않는 값은 업데이트 할 수 없음
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	// 값의 존재 여부를 판단해서 error 처리를 할 수도 있다.
	// api 에 따르면 delete 메소드는 삭제 할 값이 없다고 오류를 반환하지 않는다.
	delete(d, word)
}

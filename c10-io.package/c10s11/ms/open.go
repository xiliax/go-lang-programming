package ms

import (
	"errors"
)

func (r *MemStore) Open(n int) (err error) {
	if r.isOpened {
		return errors.New("Invalid operation, already open")
	}

	if n <= 0 {
		return errors.New("Invalid parameter, expected n > 0")
	}

	if r.capacity > 0 {
		return errors.New("Invalid operation, opened already called")
	}

	r.capacity = n
	r.isOpened = true

	return nil
}

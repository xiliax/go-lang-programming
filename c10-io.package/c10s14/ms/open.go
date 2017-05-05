package ms

import (
	"errors"
)

func (r *MemStore) Open() (err error) {
	if r.isOpened {
		return errors.New("Invalid operation, already open")
	}

	r.isOpened = true

	return nil
}

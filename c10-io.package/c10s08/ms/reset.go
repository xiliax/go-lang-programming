package ms

import (
	"errors"
)

func (r *MemStore) Reset() (err error) {
	if r.isClosed {
		return errors.New("Invalid operation for state")
	}

	r.readOffset = 0
	return nil
}

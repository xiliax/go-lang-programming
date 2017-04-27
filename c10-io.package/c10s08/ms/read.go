package ms

import (
	"errors"
)

func (r *MemStore) Read(p []byte) (n int, err error) {
	if r.isClosed {
		return 0, errors.New("Invalid operation for state")
	}

	if r.readOffset == len(r.buf) {
		return n, errors.New("EOF")
	}

	n = copy(p, r.buf[r.readOffset:])
	r.readOffset += n

	return n, nil
}

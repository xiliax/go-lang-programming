package ms

import (
	"errors"
)

func (r *MemStore) Write(p []byte) (n int, err error) {
	if r.isClosed {
		return 0, errors.New("Invalid operation for state")
	}

	l := len(p)

	temp := make([]byte, len(r.buf)+l)
	n = copy(temp, r.buf)
	n = copy(temp[n:], p)

	r.buf = temp
	return n, nil
}

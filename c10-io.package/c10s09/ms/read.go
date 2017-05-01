package ms

import (
	"io"
)

func (r *MemStore) Read(p []byte) (n int, err error) {
	if r.isClosed {
		return 0, io.ErrClosedPipe
	}

	n = copy(p, r.buf[r.readOffset:])
	r.readOffset += n

	if r.readOffset == len(r.buf) {
		return n, io.EOF
	}

	return n, nil
}

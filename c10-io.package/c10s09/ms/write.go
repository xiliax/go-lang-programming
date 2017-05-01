package ms

import "io"

func (r *MemStore) Write(p []byte) (n int, err error) {
	if r.isClosed {
		return 0, io.ErrClosedPipe
	}

	l := len(p)

	temp := make([]byte, len(r.buf)+l)
	n = copy(temp, r.buf)
	n = copy(temp[n:], p)

	r.buf = temp
	return n, nil
}

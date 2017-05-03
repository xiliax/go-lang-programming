package ms

import "io"

func (r *MemStore) Write(p []byte) (n int, err error) {
	if !r.isOpened {
		return 0, io.ErrClosedPipe
	}

	lenP := len(p)
	lenBuf := len(r.buf)

	if (lenBuf + lenP) > r.capacity {
		return 0, io.ErrShortWrite
	}

	temp := make([]byte, lenBuf+lenP)
	n = copy(temp, r.buf)
	n = copy(temp[n:], p)

	r.buf = temp
	return n, nil
}

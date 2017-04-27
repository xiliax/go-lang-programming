package ms

func (r *MemStore) Write(p []byte) (n int, err error) {
	l := len(p)

	temp := make([]byte, len(r.buf)+l)
	copy(temp, r.buf)
	n = copy(temp[n:], p)

	r.buf = temp
	return n, nil
}

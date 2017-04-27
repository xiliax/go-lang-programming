package ms

func (r MemStore) Read(p []byte) (n int, err error) {
	n = copy(p, r.buf)

	return n, nil
}

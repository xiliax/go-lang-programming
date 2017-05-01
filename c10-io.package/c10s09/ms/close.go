package ms

func (r *MemStore) Close() {
	r.isClosed = true
}

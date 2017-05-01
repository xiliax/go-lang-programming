package ms

func (r *MemStore) Close() {
	r.isOpened = false
}

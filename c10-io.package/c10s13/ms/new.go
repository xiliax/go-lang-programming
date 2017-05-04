package ms

func New(n int) *MemStore {
	if n < 0 {
		n = 0
	}

	return &MemStore{capacity: n, isOpened: true}
}

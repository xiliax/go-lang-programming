package ms

import (
	"fmt"
	"io"
)

// MemStoreOperation defines the valid operations for a 'file-like' object
type MemStoreOperation interface {
	Size() int
	IsEmpty() bool
	Write(p []byte) (n int, err error)
	io.Reader
	fmt.Stringer
}

// MemStore implenations MemStoreOperation
type MemStore struct {
	buf []byte
}

func (r MemStore) Size() int {
	return len(r.buf)
}
func (r MemStore) IsEmpty() bool {
	return 0 == r.Size()
}

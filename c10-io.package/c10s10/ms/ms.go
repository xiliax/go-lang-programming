package ms

import (
	"fmt"
	"io"
)

// MemStoreOperation defines the valid operations for a 'file-like' object
type MemStoreOperation interface {
	fmt.Stringer
	Open(int) (err error)
	Close() (err error)
	IsOpen() bool
	Write(p []byte) (n int, err error)
	io.Reader
	Size() int
	IsEmpty() bool
	Reset()
}

// MemStore implenations MemStoreOperation
type MemStore struct {
	buf        []byte
	readOffset int
	isOpened   bool
	capacity   int // maximum number of bytes stored
}

func (r MemStore) Size() int {
	return len(r.buf)
}
func (r MemStore) IsEmpty() bool {
	return 0 == r.Size()
}

func (r MemStore) IsOpen() bool {
	return r.isOpened
}

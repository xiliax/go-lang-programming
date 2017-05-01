package ms

import "fmt"

func (r MemStore) String() string {
	s := fmt.Sprintf("%T (isClosed: %v, size: %v, buf: %v)", r, r.isClosed, r.Size(), r.buf)
	return s
}

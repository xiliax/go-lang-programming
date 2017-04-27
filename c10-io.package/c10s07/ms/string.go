package ms

import "fmt"

func (r MemStore) String() string {
	return fmt.Sprintf("%T (size: %v, buf: %v)", r, r.Size(), r.buf)
}

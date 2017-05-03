package ms

import "fmt"

func (r MemStore) String() string {
	s := fmt.Sprintf("MemStroe (isOpened: %v, capacity: %v, size: %v, \n\tbuf: %v)", r.isOpened, r.capacity, r.Size(), r.buf)
	return s
}

package p

import (
	"errors"
	"fmt"
)

/*
   Read decodes a slice of bytes into a valid Person object using
   the format encoded by Person.Write()
*/
func (r Person) Read(t []byte) (n int, err error) {
	// do the work of encoding a person 'r' into bytes and store them in 'p'

	var l = len(r.Name)
	var buf = make([]byte, l+1)

	// encode r.name
	buf[0] = byte(l)
	for i := 0; i < l; i++ {
		buf[i+1] = r.Name[i]
	}

	// encode r.age
	buf = append(buf, byte(r.Age))

	// encode r.ssn
	l = len(r.Ssn)
	for i := 0; i < l; i++ {
		buf = append(buf, r.Ssn[i])
	}

	if len(t) < len(buf) {
		s := fmt.Sprintf("Buffer not big enough, required: %v, provided: %v",
			len(buf), len(t))
		return 0, errors.New(s)
	}
	n = copy(t, buf)

	return n, nil
}

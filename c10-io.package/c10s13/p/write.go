package p

import "errors"

/*
	Write encodes a Person object to bytes using:
	  Format: [
	   <name-len>:byte,
	   <bytes of name>:[]byte,
	   <age>:byte,
	   <bytes of ssn>:[11]byte
	]
*/
func (r *Person) Write(t []byte) (n int, err error) {
	if r == nil {
		return 0, errors.New("Invalid reciever")
	}

	if 0 == len(t) {
		return 0, errors.New("Invalid parameter")
	}

	nameLen := int(t[0])
	if (nameLen + 1) > len(t) {
		return 0, errors.New("Insufficient data for name")
	}
	r.Name = string(t[1 : nameLen+1])

	if (nameLen + 2) > len(t) {
		return 0, errors.New("Insufficient data for age")
	}

	r.Age = uint8(t[nameLen+1])

	buf := t[nameLen+2:]
	if ssnLen > len(buf) {
		return 0, errors.New("Insufficient data for ssn")
	}

	r.Ssn = string(buf[:ssnLen])

	return (ssnLen + 2 + nameLen), nil
}

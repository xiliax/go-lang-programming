package p

const ssnLen = 11

type Person struct {
	Name string
	Age  uint8  // 0 - 127
	Ssn  string // 999-99-9999  ( 11 byte)
}

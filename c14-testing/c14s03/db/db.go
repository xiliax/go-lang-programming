package db

type Id uint64
type Db interface {
	Open() (err error)
	Close() (err error)
	Insert(d []byte) (id Id, err error)
	Get(d []byte, id Id) (n int, err error)
}

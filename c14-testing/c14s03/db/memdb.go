package db

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/another/c14-testing/c14s01/db"
)

const (
	maxCapacity = 1024 * 1024
	invalidId   = 0
)

var (
	errTooBig     = errors.New(fmt.Sprintf("specified capacity greater than %v", maxCapacity))
	errDbClose    = errors.New("operation not premitted on close db")
	errInsufSpace = errors.New("insufficient space to save data")
)

type MemDb struct {
	buf      bytes.Buffer
	open     bool
	capacity uint64
	idCh     chan int
}

func New(capacity uint64) (db *MemDb, err error) {
	if capacity > maxCapacity {
		return nil, errTooBig
	}

	db = &MemDb{capacity: capacity}
	db.idCh = make(chan int, 1)
	db.idCh <- 0
	return
}

func (db *MemDb) Open() (err error) {
	db.open = true
}

func (db *MemDb) Close() (err error) {
	db.open = false
}

func (db *MemDb) Insert(d []byte) (id db.Id, err error) {
	if !db.open {
		return invalidId, errDbClose
	}

	if db.buf.Len()+len(d) > db.capacity {
		return invalidId, errInsufSpace
	}

	db.buf.Write(d)
	id = db.nextId()
	return
}

func (db *MemDb) Get(d []byte, id db.Id) (n int, err error) {
	panic("not implemented")
}

func (db *MemDb) nextId() (id Id) {
	i := <-db.idCh + 1
	db.idCh <- i
	id = Id(i)
	return
}

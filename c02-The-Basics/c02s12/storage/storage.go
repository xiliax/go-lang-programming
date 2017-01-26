package storage

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

func init() {
	log.Println("Package 'storage' being initialized")
	rand.Seed(time.Now().Unix())
}

// StoreResult takes and int and writes it to the storage engine
func StoreResult(r int) error {
	db, err := getStorageEngine()
	if nil != err {
		return err
	}
	*db = r
	return nil
}

func getStorageEngine() (*int, error) {
	i := rand.Intn(2)
	var d = 0
	if i == 0 {
		return &d, nil
	}

	return nil, errors.New("failure to get storage engine")
}

// GetData returns data if it is available, otherwise it returns an error
func GetData() (*string, error) {
	i := rand.Intn(3)

	var d string

	switch i {
	case 0:
		d = "this is a lot of data"
	case 1:
		d = "this is even more data than before"
	default:
		return nil, errors.New("there is no data")
	}
	return &d, nil
}

/*
Awesome application reads data from storage, runs complex computation,
and writes results back to storage.
*/
package main

import (
	"fmt"
	"log"

	_ "github.com/another/c02-The-Basics/c02s12/silly"

	"github.com/another/c02-The-Basics/c02s12/compute"
	"github.com/another/c02-The-Basics/c02s12/storage"
)

func init() {
}

func main() {
	fmt.Println("My Awesome Application")

	// TODO 1 - Get data from storage
	data, err := storage.GetData()
	if nil != err {
		log.Fatalln(err)
	}
	// TODO 2 - Run complex computation
	result := compute.Calc(data)
	// TODO 3 - Save result to storage
	if err = storage.StoreResult(result); nil != err {
		log.Fatalln(err)
	}

	log.Println("Awesome finished successfully")
}

package main

import (
	"io/ioutil"
	"log"
)

const (
	wpFilename = "welcome.html"
	hpFilename = "home.html"
)

var (
	homePage    string
	welcomePage string
)

func init() {
	// load pages from file
	wpBytes, err := ioutil.ReadFile(wpFilename)
	if nil != err {
		log.Fatalln(err)
	}

	welcomePage = string(wpBytes)

	hpBytes, err := ioutil.ReadFile(hpFilename)
	if nil != err {
		log.Fatalln(err)
	}
	homePage = string(hpBytes)
}

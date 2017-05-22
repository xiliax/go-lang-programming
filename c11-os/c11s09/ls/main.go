package main

import (
	"fmt"
	"os"
)

/*
For each operand that names a file of a type other than directory, ls dis-
     plays its name as well as any requested, associated information.  For each
     operand that names a file of type directory, ls displays the names of files
     contained within that directory, as well as any requested, associated infor-
     mation.

     If no operands are given, the contents of the current directory are dis-
     played.  If more than one operand is given, non-directory operands are dis-
     played first;
*/
var longListing = false

func main() {
	offset := 1 // skip program name
	if 2 <= len(os.Args) && "-l" == os.Args[offset] {
		longListing = true
		offset++
	}

	files := os.Args[offset:]
	if 0 == len(files) {
		files = []string{"."} // show current directory
	}

	if longListing {
		showLongListing(files)
		return
	}

	showShortListing(files)
}

func showShortListing(files []string) {
	var noFilesList []string
	var filesList []string
	var dirListing []string
	for _, f := range files {
		fi, err := os.Stat(f)
		if nil != err {
			s := fmt.Sprintf("ls: %v: no file or directory", f)
			noFilesList = append(noFilesList, s)
			continue
		}

		if !fi.IsDir() {
			filesList = append(filesList, f)
			continue
		}

		// get list of files in directory
		dirListing = addShortDirListing(dirListing, f)
	}

	for _, s := range noFilesList {
		fmt.Println(s)
	}
	for _, s := range filesList {
		fmt.Println(s)
	}
	for _, s := range dirListing {
		fmt.Println(s)
	}

}

func addShortDirListing(listing []string, f string) []string {
	dir, err := os.Open(f)
	if nil != err {
		return listing
	}

	fileNames, err := dir.Readdirnames(0)
	if nil != err {
		return listing
	}

	listing = append(listing, "\n"+f+":")
	for _, d := range fileNames {
		listing = append(listing, d)
	}

	return listing
}

func addLongDirListing(listing []string, f string) []string {
	dir, err := os.Open(f)
	if nil != err {
		return listing
	}

	fis, err := dir.Readdir(0)
	if nil != err {
		return listing
	}

	listing = append(listing, "\n"+f+":")
	s := ""
	for _, fi := range fis {
		size := calSize(fi.Size())
		perm := fi.Mode().Perm()

		s = fmt.Sprintf("%s %v %s", perm, size, fi.Name())
		listing = append(listing, s)
	}

	return listing
}

func showLongListing(files []string) {
	var noFilesList []string
	var filesList []string
	var dirListing []string
	for _, f := range files {
		fi, err := os.Stat(f)
		if nil != err {
			s := fmt.Sprintf("ls: %v: no file or directory", f)
			noFilesList = append(noFilesList, s)
			continue
		}

		if !fi.IsDir() {
			size := calSize(fi.Size())
			perm := fi.Mode().Perm()
			s := fmt.Sprintf("%s %v %s", perm, size, f)
			filesList = append(filesList, s)
			continue
		}

		// get list of files in directory
		dirListing = addLongDirListing(dirListing, f)
	}

	for _, s := range noFilesList {
		fmt.Println(s)
	}
	for _, s := range filesList {
		fmt.Println(s)
	}
	for _, s := range dirListing {
		fmt.Println(s)
	}
}

func calSize(i int64) string {
	s := float64(i)
	unit := "B" // bytes

	if 1.0 < (s / 1024) {
		s = s / 1024
		unit = "K"
	}

	if 1.0 < (s / 1024) {
		s = s / 1024
		unit = "M"
	}

	if 1.0 < (s / 1024) {
		s = s / 1024
		unit = "G"
	}

	if 1.0 < (s / 1024) {
		s = s / 1024
		unit = "T"
	}

	return fmt.Sprintf("%6.2f%v", s, unit)
}

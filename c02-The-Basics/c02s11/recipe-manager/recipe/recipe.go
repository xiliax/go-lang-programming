package recipe

import "fmt"

var db = make(map[string]string)

// Add is a function to add a recipe to our recipe db
func Add(name string) {
	fmt.Println("Adding new recipe '" + name + "'")
	db[name] = name
}

// Remove is a function to remove a recipe form our recipe db
func Remove(name string) {
	if _, present := db[name]; present {
		fmt.Println("Removing recipe '" + name + "'")
		delete(db, name)
	}
}

// Count returns the number of recipies in db
func Count() int {
	return len(db)
}

package main

import (
	"fmt"

	"github.com/another/c02-The-Basics/c02s08/recipe-manager/recipe"
	"github.com/another/c02-The-Basics/c02s08/security"
)

const username = "Verrol"

func main() {
	if !security.IsAuthenticated(username) {
		fmt.Printf("%s is not authenticated, logging\n", username)
		security.Login(username)
	}

	recipe.Add("Pizza")
	recipe.Add("Cake")
	recipe.Remove("Sushi")
	recipe.Remove("Cake")

	security.Logout()

	if security.IsAuthenticated(username) {
		fmt.Println("ERROR - User still authentiated")
	} else {
		fmt.Println("Great, user successfully logged out")
	}
}

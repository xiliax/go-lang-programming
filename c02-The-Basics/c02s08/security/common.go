package security

// package variable, private
var currentUser string = ""

// IsAuthenticated checks if this user is authenticated
func IsAuthenticated(username string) bool {
	return username == currentUser
}

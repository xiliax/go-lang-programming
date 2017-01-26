package compute

import "log"

func init() {
	log.Println("Package 'compute' being initialized")
}

func Calc(d *string) int {
	return len(*d)
}

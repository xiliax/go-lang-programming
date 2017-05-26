package f

import "os"

func Print(s ...interface{}) {
	Fprint(os.Stdout, s...)
}

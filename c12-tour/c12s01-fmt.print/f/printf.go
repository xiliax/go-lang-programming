package f

import "os"

func Printf(format string, values ...interface{}) {
	s := Sprintf(format, values...)
	Fprint(os.Stdout, s)
}

package f

import "io"

func Fprintf(w io.Writer, format string, values ...interface{}) {
	s := Sprintf(format, values...)
	Fprint(w, s)
}

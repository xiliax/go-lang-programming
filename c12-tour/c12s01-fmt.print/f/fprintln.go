package f

import "io"

func Fprintln(w io.Writer, s ...interface{}) {
	s = append(s, "\n")
	Fprint(w, s...)
}

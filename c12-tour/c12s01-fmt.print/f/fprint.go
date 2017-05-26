package f

import "io"

func Fprint(w io.Writer, s ...interface{}) {
	t := Sprintf("%s", s...)
	io.WriteString(w, t)
}

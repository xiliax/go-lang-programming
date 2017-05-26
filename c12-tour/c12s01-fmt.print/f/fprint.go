package f

import "io"

func Fprint(w io.Writer, s ...interface{}) {
	for _, a := range s {
		switch i := a.(type) {
		case string:
			t := Sprintf("%s", string(i))
			io.WriteString(w, t)
		case int:
			io.WriteString(w, itoa(int(i)))
		case bool:
			v := "false"
			if bool(i) {
				v = "true"
			}
			io.WriteString(w, v)
		case float64:
			io.WriteString(w, "[float64]")
		}
	}
}

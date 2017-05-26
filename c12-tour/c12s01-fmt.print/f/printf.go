package f

func Printf(format string, values ...interface{}) {
	s := Sprintf(format, values...)
	Print(s)
}

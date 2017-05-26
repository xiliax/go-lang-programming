package f

func Println(s ...interface{}) {
	s = append(s, "\n")
	Print(s...)
}

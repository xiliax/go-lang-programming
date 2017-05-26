package f


var lookup = []string{
	"0", "1", "2", "3", "4", "5",
	"6", "7", "8", "9",
}

func itoa(i int) (s string) {
	if 0 > i {
		s = "-"
		i = -1 * i
	}

	t := ""
	for 0 < i {
		r := i % 10
		i = i / 10
		t = lookup[r] + t
	}
	return (s + t)
}

package f

func Sprintf(format string, values ...interface{}) (r string) {
	offset := 0
	for i := 0; i < len(format); i++ {
		k := format[i]
		switch k {
		case '%':
			// do special converstion
			switch format[i+1] {
			case 's':
				r = r + values[offset].(string)
			case 'd':
				r = r + itoa(values[offset].(int))
			case '.':
			// consume number or char after '.'
			default:
				r = r + "(unhandled)"
			}

			i++ // skip format
			offset++

		default:
			r = r + string(k)
		}
	}
	return
}

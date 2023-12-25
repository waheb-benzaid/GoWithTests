package iteration

func Repeat(char string, repeatedTimes int) string {
	repeated := ""
	for i := 0; i < repeatedTimes; i++ {
		repeated += char
	}

	return repeated
}
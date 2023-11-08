package iteration

func Repeat(char string, repeatCount int) (res string) {
	for i := 0; i < repeatCount; i++ {
		res += char
	}

	return
}

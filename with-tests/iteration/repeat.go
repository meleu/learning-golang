package iteration

func Repeat(character string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}

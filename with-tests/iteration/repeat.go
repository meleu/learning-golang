package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
	}

	return repeated
}

func Repeat2(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}

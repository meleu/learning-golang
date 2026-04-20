package iteration

import "strings"

const repeatCount = 5

func RepeatNaive(character string) string {
	var repeated string
	for range repeatCount {
		repeated = repeated + character
	}
	return repeated
}

func Repeat(character string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

package loops

import "strings"

// Repeat returns character repeated 5 times.
func Repeat(input string, times int) (output string) {
	if times <= 0 {
		output = ""
	}

	for i := 0; i < times; i++ {
		output += input
	}
	return output
}

func Upper(input string, times int) (output string) {
	for i:= 0; i < times; i++ {
		output += strings.ToUpper(input)
	}
	return output
}

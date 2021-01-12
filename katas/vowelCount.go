package vowelCount

import "fmt"

//https://www.codewars.com/kata/54ff3102c1bad923760001f3/go

func main() {
	for i := 0; i < len(str); i++ {
		char := (string([]rune(str)[i]))
		switch char {
		case "a":
			count++
		case "e":
			count++
		case "i":
			count++
		case "o":
			count++
		case "u":
			count++
		default:
		}
	}
	fmt.Println(count)
}

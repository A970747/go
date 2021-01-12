package vowelCount

import "fmt"

//https://www.codewars.com/kata/54ff3102c1bad923760001f3/go

func vowelCount(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	fmt.Println(count)
}

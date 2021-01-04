package main

func Sum(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
}

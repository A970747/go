package sum

import "fmt"

func Sum(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllSlices(slices ...[]int) {
	var sum []int
	var innerSum int
	for _, slice := range slices {
		innerSum = 0
		for _, value := range slice {
			innerSum += value
		}
		sum = append(sum, innerSum)
		fmt.Println(sum)
	}
}

func main() {
}

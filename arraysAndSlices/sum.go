package sum

import (
	"fmt"
)

//Sum sums a single slice
func Sum(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//AllSlices Sums any number of slices into a separate slice
func AllSlices(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return sums
}

//AllTails Adds the tails of all slices to a new slice
func AllTails(slices ...[]int) (sums []int) {
	fmt.Println(slices[0])
	for _, slice := range slices {
		tail := slice[0:]
		sums = append(sums, Sum(tail))
	}
	return sums
}

func main() {
}

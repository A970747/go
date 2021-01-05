package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertMessage := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("sum collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertMessage(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertMessage(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	assertMessage := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("multiple slices return slice with sum of sub-slices", func(t *testing.T) {
		got := SumAllSlices([]int{1, 2}, []int{1, 3, 9})
		want := []int{3, 13}
		assertMessage(t, got, want)
	})
}

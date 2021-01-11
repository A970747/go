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

	t.Run("sum collection of some numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
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
		got := AllSlices([]int{1, 2}, []int{6, 9})
		want := []int{3, 15}
		assertMessage(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assertMessage := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("sum tails of passed slices", func(t *testing.T) {
		got := AllTails([]int{}, []int{6, 9, 11})
		want := []int{0, 26}
		assertMessage(t, got, want)
	})
}

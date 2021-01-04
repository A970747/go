package loops

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("got %q want %q", expected, repeated)
		}
	}

	t.Run("test allow user to input repeat", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("test return nothing if repeat input 0 or less", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := ""
		assertCorrectMessage(t, expected, repeated)
	})
}

func TestUpper(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("got %q want %q", expected, repeated)
		}
	}

	t.Run("all input characters uppercase & repeated n times", func(t *testing.T) {
		repeatedUpper := Upper("a", 5)
		expectedUpper := "AAAAA"

		assertCorrectMessage(t, expectedUpper, repeatedUpper)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Upper("a", 5)
	}
}

func ExampleRepeat() {
	output := Repeat("a", 5)
	fmt.Println(output)
	// Output: aaaaa
}

func ExampleUpper() {
	output := Upper("a", 5)
	fmt.Println(output)
	// Output: AAAAA
}

package main

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

type SpySleeper struct {
	Calls int
}

type CountdownOperationsSpy struct {
	Calls []string
}

func TestCountdown(t *testing.T) {

	t.Run("base test, no ordering", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("not enoug calls on sleeper, want 4 got %d", spySleeper.Calls)
		}
	})

	t.Run("checks Sleep executes in correct order", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

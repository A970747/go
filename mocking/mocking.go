package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

//DefaultSleeper default struct so we can mock in test
type DefaultSleeper struct{}

//Sleeper interface so we can pass a different implementation of Sleep in test
type Sleeper interface {
	Sleep()
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

//Countdown times out outputs
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

//Sleep default sleep function so we can mock this
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

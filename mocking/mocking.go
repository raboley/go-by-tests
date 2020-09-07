package mocking

import (
	"fmt"
	"io"
	"time"
)

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	_, _ = fmt.Fprint(writer, finalWord)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

const sleep = "sleep"
const write = "write"

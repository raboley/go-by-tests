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

// ConfigurableSleeper takes in a duration of time and a sleep method.
// Using these two fields one can configure something that executes a Sleep command
// to use the duration for that sleep command.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration: duration, sleep: sleep}
}

const finalWord = "Go!"
const countdownStart = 3

const sleep = "sleep"
const write = "write"

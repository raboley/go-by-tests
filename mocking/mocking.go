package mocking

import (
	"fmt"
	"io"
	"time"
)

// Countdown will iterate downward from const countdownStart to 0 executing the sleeper's Sleep() method
// in-between each iteration, and printing a line to the writer. It will finish by performing one last
// sleep then printing out the const finalWord.
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

// Sleep will perform the private sleep method passing in the duration
// private field as the only argument.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// NewConfigurableSleeper initializes a *ConfigurableSleeper setting the private fields duration and sleep.
func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration: duration, sleep: sleep}
}

const finalWord = "Go!"
const countdownStart = 3

const sleep = "sleep"
const write = "write"

package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &CountdownOperationsSpy{}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
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

func TestConfigurableSleeper_Sleep(t *testing.T) {
	sleepTime := 10 * time.Second

	spy := &TimeSpy{}
	sleeper := ConfigurableSleeper{sleepTime, spy.Sleep}
	sleeper.Sleep(sleepTime)

	got := spy.durationSlept
	want := sleepTime

	if want != got {
		t.Errorf("wanted sleep duration to be %v but got sleep duration of %v", want, got)
	}

}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(_ []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type TimeSpy struct {
	durationSlept time.Duration
}

func (s *TimeSpy) Sleep(duration time.Duration) {
	s.durationSlept += duration
}

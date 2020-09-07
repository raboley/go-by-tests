package mocking

import (
	"bytes"
	"os"
	"reflect"
	"testing"
	"time"
)

func ExampleCountdown() {
	sleeper := NewConfigurableSleeper(0*time.Second, time.Sleep)
	Countdown(os.Stdout, sleeper)
	// Output: 3
	//2
	//1
	//Go!
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &spyCountDownOperations{}
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
		spySleepPrinter := &spyCountDownOperations{}
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

func ExampleConfigurableSleeper_Sleep() {
	sleeper := NewConfigurableSleeper(0*time.Second, time.Sleep)
	sleeper.Sleep()
}

func TestConfigurableSleeper_Sleep(t *testing.T) {
	sleepTime := 10 * time.Second

	spy := &spyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spy.Sleep}
	sleeper.Sleep()

	got := spy.durationSlept
	want := sleepTime

	if want != got {
		t.Errorf("wanted sleep duration to be %v but got sleep duration of %v", want, got)
	}

}

// spyCountDownOperations is a test spy that records what should have
// been written to the writer, and how many calls to the Sleep() method
// were made.
type spyCountDownOperations struct {
	Calls []string
}

func (s *spyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *spyCountDownOperations) Write(_ []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// spyTime is a test spy to count the duration
// of time that was supposed to be spent sleeping.
type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.durationSlept += duration
}

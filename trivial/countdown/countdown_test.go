package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 and go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("Order of operations", func(t *testing.T) {
		spyCountdownOperations := &SpyCountdownOperations{}
		Countdown(spyCountdownOperations, spyCountdownOperations)
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

		if !reflect.DeepEqual(want, spyCountdownOperations.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyCountdownOperations.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 4 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

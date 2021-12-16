package countdown

import "time"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "Sleep"

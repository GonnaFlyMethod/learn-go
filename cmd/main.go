package main

import (
	"learn_go/trivial/countdown"
	"os"
	"time"
)

func main() {
	sleeper := countdown.ConfigurableSleeper{time.Second, time.Sleep}
	countdown.Countdown(os.Stdout, &sleeper)
}

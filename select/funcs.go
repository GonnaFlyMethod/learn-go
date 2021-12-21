package _select

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondsTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

//func Racer(a, b string) (winner string) {
//	aDuration := measureResponseTime(a)
//	bDuration := measureResponseTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//	return b
//}
//
//func measureResponseTime(url string) time.Duration {
//	startB := time.Now()
//	http.Get(url)
//	return time.Since(startB)
//}
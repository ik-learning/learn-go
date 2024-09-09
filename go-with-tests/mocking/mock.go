package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprintf(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
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
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(d time.Duration)
}

func (s ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

type SpyTime struct {
	duration time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.duration = duration
}

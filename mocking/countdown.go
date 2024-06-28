package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
  Sleep()
}

type DefaultSleeper struct{}

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

func (d *DefaultSleeper) Sleep() {
  time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    fmt.Fprintln(out, i)
    sleeper.Sleep()
  }
  fmt.Fprint(out, finalWord)
}

func main() {
  sleeper := &DefaultSleeper{}
  Countdown(os.Stdout, sleeper)
}

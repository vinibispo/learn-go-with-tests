package main

import (
	"bytes"
	"slices"
	"testing"
)

func TestCountdown(t *testing.T) {
  t.Run("prints 3 to Go!", func(t *testing.T) {
    buffer := &bytes.Buffer{}
    Countdown(buffer, &SpyCountdownOperations{})
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
    spySleepPrinter := &SpyCountdownOperations{}
    Countdown(spySleepPrinter, spySleepPrinter)

    want := []string{
      write,
      sleep,
      write,
      sleep,
      write,
      sleep,
      write,
    }

    if !slices.Equal(want, spySleepPrinter.Calls) {
      t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
    }
  })
}

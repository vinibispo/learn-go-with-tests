package main

import (
	"slices"
	"testing"
)

func TestCountdown(t *testing.T) {
  t.Run("sleep before print", func(t *testing.T) {
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

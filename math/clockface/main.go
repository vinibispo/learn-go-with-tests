package main

import (
	"os"
	"time"

	clockface "github.com/vinibispo/learn-go-with-tests/math"
)

func main() {
	t := time.Now()
  clockface.SVGWriter(os.Stdout, t)
}

package main

import (
	"os"
	"time"

	"github.com/vinibispo/learn-go-with-tests/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}

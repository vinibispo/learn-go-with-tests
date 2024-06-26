package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
  repeated := Repeat("a")
  expected := "aaaa"

  if repeated != expected {
    t.Errorf("expected %q got %q", expected, repeated)
  }
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a")
  }
}

// Repeat takes one character and repeats four times
func ExampleRepeat() {
  result := Repeat("a")
  fmt.Println(result)
}

package integers

import (
	"fmt"
	"testing"
)


func TestAdder(t *testing.T) {
  sum := Add(2, 2)
  expected := 4

  if sum != expected {
    t.Errorf("expected '%d' but got '%d'", expected, sum)
  }
}

// Add takes two integers and returns the sum of them.
func ExampleAdd() {
  sum := Add(1, 5)
  fmt.Println(sum)
}

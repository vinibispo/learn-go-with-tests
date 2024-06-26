package iteration

import "testing"

func TestRepeat(t *testing.T) {
  repeated := Repeat("a")
  expected := "aaaa"

  if repeated != expected {
    t.Errorf("expected %q got %q", expected, repeated)
  }
}

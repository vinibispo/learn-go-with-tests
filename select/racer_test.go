package racer

import "testing"

func TestRacer(t *testing.T) {
  slowUrl := "https://facebook.com"
  fastUrl := "http://quii.dev"
  want := fastUrl
  got := Racer(slowUrl, fastUrl)

  if got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}

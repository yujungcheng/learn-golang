package greeting

import (
  "testing"
)

func TestGreeting(t *testing.T) {
  got := Greeting("George")
  want := "Hello George"
  if got != want {
    t.Fatalf("Greeting: Expected %q, got %q", want, got)
  }
}

package utils

import (
    "testing"
)

// TestAddNumber calls addNumber with a two numbers, checking
// for a valid return value.
func TestAddNumbers(t *testing.T) {
    expectedAnswer := 10
    want := expectedAnswer
    msg := AddNumbers(3,7)
    if want != msg  {
        t.Fatalf(`addNumbers(3, 7) = %q, want match for %#q, nil`, msg, want)
    }
}
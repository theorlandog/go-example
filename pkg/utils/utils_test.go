package utils

import (
    "testing"
)

// TestAddNumber calls utils.addNumber with a two numbers, checking
// for a valid return value.
func TestAddNumber(t *testing.T) {
    expectedAnswer := 10
    want := expectedAnswer
    msg := addNumber(3,7)
    if want != msg  {
        t.Fatalf(`addNumber(3, 7) = %q, want match for %#q, nil`, msg, want)
    }
}
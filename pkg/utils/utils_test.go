package utils

import (
    "testing"
)

// TestAddNumber calls addNumber with a two numbers, checking
// for a valid return value.
func TestAddNumbers(t *testing.T) {
    expectedAnswer := 10
    answer := AddNumbers(3,7)
    if expectedAnswer != answer  {
        t.Fatalf(`addNumbers(3, 7) = %q, want match for %#q, nil`, answer, expectedAnswer)
    }
}
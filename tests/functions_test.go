package tests

import (
	"goTutorial/functions"
	"testing"
)

// test functions must be declared as public and start with "test" keyword.
func TestSumFunction(t *testing.T) {
	sum := functions.Sum(3, 5)
	if sum != 8 {
		t.Errorf("sum failed")
	}
}

func TestSubFunction(t *testing.T) {
	sub := functions.Sub(3, 5)
	if sub != -2 {
		t.Errorf("sub failed")
	}
}

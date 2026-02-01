package main

import "testing"

func TestFactorial(t *testing.T) {
	got := factorial(5)
	want := 120

	if got != want {
		t.Errorf("Factorial error: Expected %d, returned %d.", want, got)
	}
}

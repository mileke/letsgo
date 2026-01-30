package main

import (
	"log"
	"testing"
)

func TestOperationAdd(t *testing.T) {
	got := doOperation(5, 4, add)
	want := 9

	if got != want {
		log.Fatalf("Addition error: Got %d, want %d.", got, want)
	}
}
func TestOperationSubtract(t *testing.T) {
	got := doOperation(5, 4, subtract)
	want := 2

	if got != want {
		log.Fatalf("Subtraction error: Got %d, want %d.", got, want)
	}
}

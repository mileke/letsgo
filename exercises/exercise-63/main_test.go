package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)

	if total != 9 {
		t.Errorf("The test has failed. Expected %d, but got %d", 9, total)
	}
}

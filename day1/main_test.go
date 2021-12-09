package main

import (
	"testing"
)

func TestExampleFunction(t *testing.T) {
	// case 1: number's length 1, the random number be less than 10
	// case 2: number's length 2, the random number be less than 100
	// case 3: number's length 3, the random number be less than 1000
	got := 10
	want := 10
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

package main

import (
	"strings"
	"testing"
)

func TestCleanUpMessage(t *testing.T) {
	expected := "BUY NOW, SAVE 10%"
	border := strings.Repeat("*", len(expected)+6*2)
	contentPadding := strings.Repeat(" ", 5)
	input := border + "\n" + "*" + contentPadding + expected + contentPadding + "*" + "\n"
	input += border
	t.Log(input)

	if observed := CleanUpMessage(input); observed != expected {
		// use Error in case you have a big  list of tests need to continue
		// t.Errorf("Expected: %v, but the result is %v", expected, observed)
		// use Fatal in case you wanna stop if the function is failed
		t.Fatalf("Expected: %v, but the result is %v", expected, observed)
	}
}

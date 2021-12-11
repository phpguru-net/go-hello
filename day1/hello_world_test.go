package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {

	expected := "Hello, World!"
	if observed := HelloWorld(); observed != expected {
		// use Error in case you have a big  list of tests need to continue
		// t.Errorf("Expected: %v, but the result is %v", expected, observed)
		// use Fatal in case you wanna stop if the function is failed
		t.Fatalf("Expected: %v, but the result is %v", expected, observed)
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		HelloWorld()
	}
}

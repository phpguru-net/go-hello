package main

import (
	"testing"
)

func TestInterestRateNegativeBalance(t *testing.T) {
	tt := struct {
		name    string
		balance float64
		want    float32
	}{

		name:    "Minimal second interest rate",
		balance: 1000.0,
		want:    float32(1.621),
	}

	got := InterestRate(tt.balance)

	if got != tt.want {
		t.Fatalf("Balance is %v. Expected Interest Rate %v but got %v", tt.balance, tt.want, got)
	}
}
func TestInterestRate1KRangeBalance(t *testing.T) {
	balance := float64(500)
	expected := float32(0.5)

	if got := InterestRate(balance); got != expected {
		t.Fatalf("Balance is %v. Expected Interest Rate %v but got %v", balance, expected, got)
	}
}

func TestInterestRate5KRangeBalance(t *testing.T) {
	balance := float64(1500)
	expected := float32(1.612)

	if got := InterestRate(balance); got != expected {
		t.Fatalf("Balance is %v. Expected Interest Rate %v but got %v", balance, expected, got)
	}
}

func TestInterestRateMoreThan5KRangeBalance(t *testing.T) {
	balance := float64(10000)
	expected := float32(2.475)

	if got := InterestRate(balance); got != expected {
		t.Fatalf("Balance is %v. Expected Interest Rate %v but got %v", balance, expected, got)
	}
}

func TestInterestNegativeBalance(t *testing.T) {
	balance := float64(-1000)
	expected := float64(-1003.213)

	if got := Interest(balance); got != expected {
		t.Fatalf("Balance is %v. Expected Interest Rate %v but got %v", balance, expected, got)
	}
}

func TestYearsBeforeDesiredBalance(t *testing.T) {
	balance := float64(2345.67)
	targetBalance := float64(12345.6789)
	expected := 85
	if got := YearsBeforeDesiredBalance(balance, targetBalance); got != expected {
		t.Fatalf("Target Year %v but got %v", expected, got)
	}
}

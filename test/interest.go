package main

import (
	"math"
)

func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return float32(-3.213)
	case balance >= 0 && balance < 1000:
		return float32(0.5)
	case balance >= 1000 && balance < 5000:
		return float32(1.621)
	default:
		return float32(2.475)
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

func YearsBeforeDesiredBalanceWithStaticRate(balance, targetBalance float64) int {
	// xn = x0 * (1+r)^n
	// xn - x0 = (1+r)^n
	// n = log(1+r)(xn-xo)
	// n = log10(xn/xo)/ log10(1+r)
	log10Interest := math.Log10(float64(1 + InterestRate(balance)/100.0))
	log10Balance := math.Log10(targetBalance) - math.Log10(balance)
	return int(math.Ceil(log10Balance / log10Interest))
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// we have 3 ranges:
	// [0-1000}: r1
	// [1000-5000}: r2
	// [5000-]: r3
	years := 0
	interationBalance := balance
	for interationBalance < targetBalance {
		years += 1
		interationBalance = AnnualBalanceUpdate(interationBalance)
	}

	return years
}

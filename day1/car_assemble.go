package main

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	}
	if speed >= 1 && speed <= 4 {
		return 1.0
	}
	if speed >= 5 && speed <= 8 {
		return 0.9
	}
	// otherwise
	return 0.77
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	if speed == 0 {
		return 0.0
	}
	const baseProductionRate = 221
	return float64(baseProductionRate) * float64(speed) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	var productionRatePerHour float64 = CalculateProductionRatePerHour(speed)
	if productionRatePerHour > limit {
		return limit
	}
	// otherwise
	return productionRatePerHour
}

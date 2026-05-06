package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate/100.0)*float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successPerHour:= (successRate/100.0)*float64(productionRate)
    return int(successPerHour/60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens:= carsCount/10
    rem := carsCount%10
    if carsCount>=10{
        return uint(tens*95000 + rem * 10000)
    }
    return uint(carsCount * 10000) 
}

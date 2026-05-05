// Package weather contains all the functions used for weather prediction.
package weather

var (
    // CurrentCondition variable stores the current wether condintion.
	CurrentCondition string 
    // CurrentLocation variable stores the current location.
	CurrentLocation  string 
)

// Forecast makes forecass for a city using its current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

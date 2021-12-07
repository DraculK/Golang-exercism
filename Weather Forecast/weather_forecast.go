// Package weather is a package used to inform the current location and weather condition of someplace.
package weather

// CurrentCondition informs the user the weather condition at the time.
var CurrentCondition string

// CurrentLocation informs the user his location.
var CurrentLocation string

// Forecast function prints the Current Location and the Current Condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
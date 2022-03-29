// Package weather is set of feature about weather.
package weather

// CurrentCondition is current condition.
var CurrentCondition string

// CurrentLocation is current location.
var CurrentLocation string

// Forecast return current location's current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

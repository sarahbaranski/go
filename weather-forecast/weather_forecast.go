// Package weather provides the current forcast
// based on the location of city and condition.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location used in the forcast.
var CurrentLocation string

// Forecast returns a string value equal to the current city minus the current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

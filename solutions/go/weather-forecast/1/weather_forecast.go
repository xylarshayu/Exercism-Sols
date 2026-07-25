// Package weather is for declaring
// weather forecasts.
package weather

var (
	// CurrentCondition states the current
	// weather condition.
	CurrentCondition string
	// CurrentLocation states the current
	// location.
	CurrentLocation string
)

// Forecast returns a declaration of
// the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

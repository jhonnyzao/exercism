// Package space provides functions to calculate
// someone's age in different planets of the solar system
package space

// Age expects a number of seconds
// and a to calculate someone's age
func Age(seconds float64, planet string) float64 {
	const earthYearInSeconds = 31557600

	proportions := map[string]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return seconds / earthYearInSeconds / proportions[planet]
}

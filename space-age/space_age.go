package space

import "strings"

var earthAgeSec = 31557600

var earthYears = map[string]float64{
	"earth":   1.0,
	"jupiter": 11.862615,
	"mars":    1.8808158,
	"mercury": 0.2408467,
	"neptune": 164.79132,
	"saturn":  29.447498,
	"uranus":  84.016846,
	"venus":   0.61519726,
}

// Age return someone's age in years on the planet
func Age(sec float64, planet string) float64 {
	planetAgeSec := earthYears[strings.ToLower(planet)] * float64(earthAgeSec)
	return sec / planetAgeSec
}

package space

func Age(time float64, planet string) float64 {

	planetsPeriod := map[string]float64{
		"Earth": 1,
		"Mercury": 0.2408467,
		"Venus": 0.61519726,
		"Mars": 1.8808158,
		"Jupiter": 11.862615,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}
	earthSecondsOrbit := 31557600

	planetPeriod, ok := planetsPeriod[planet]

	if ok {
		return time / (float64(earthSecondsOrbit) * planetPeriod)
	} else {
		return 0
	}
}

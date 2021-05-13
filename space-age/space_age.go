package space

type Planet string

func Age(seconds float64, pt Planet) (year float64) {
	switch pt{
	case "Earth":
		year = seconds / (365.25 * 24 * 60 * 60)
	case "Mercury":
		year = seconds / (365.25 * 0.2408467 * 24 * 60 * 60)
	case "Venus":
		year = seconds / (365.25 * 0.61519726 * 24 * 60 * 60)
	case "Mars":
		year = seconds / (365.25 * 1.8808158 * 24 * 60 * 60)
	case "Jupiter":
		year = seconds / (365.25 * 11.862615 * 24 * 60 * 60)
	case "Saturn":
		year = seconds / (365.25 * 29.447498 * 24 * 60 * 60)
	case "Uranus":
		year = seconds / (365.25 * 84.016846 * 24 * 60 * 60)
	case "Neptune":
		year = seconds / (365.25 * 164.79132 * 24 * 60 * 60)
	}
	return year
}
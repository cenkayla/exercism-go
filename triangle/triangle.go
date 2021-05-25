package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
    // Pick values for the following identifiers used by the test program.
    NaT = "Üçgen değil" // not a triangle
    Equ = "Eşkenar üçgen" // equilateral
    Iso = "İkizkenar üçgen" // isosceles
    Sca = "Çeşitkenar üçgen" // scalene
)

func KindFromSides(a, b, c float64) (k Kind) {

	if a <= 0 || b <= 0 || c <= 0 {
		k = "Üçgen değil"
	}else if a + b < c || b + c < a || a + c < b {
		k = "Üçgen değil"
	}else if a != a || b!=b || c != c {
		k = "Üçgen değil"
	}else if math.IsInf(a,1) || math.IsInf(b,1) || math.IsInf(c,1) {
		k = "Üçgen değil"
	}else if a == b && b == c  {
		k = "Eşkenar üçgen"
	}else if a == b  || b == c || a == c {
		k = "İkizkenar üçgen"
	}else if a != b && b != c {
		k = "Çeşitkenar üçgen"
	}

	return k
}

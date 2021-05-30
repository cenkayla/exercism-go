package prime

import "math"

func Nth(n int) (int, bool) {
	num := 1
	if n <= 0 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}
	 for {
		num += 2
		if IsPrime(num) {
			n--
		} 
		if n == 1 {
			return num, true
		}
	}
}
func IsPrime(n int) bool {
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
} 


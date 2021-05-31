package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(num int) bool {

	var sum int
	s := strconv.Itoa(num)
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return false
	}
	for _, v := range s {
		sum += int(math.Pow(float64(int(v - '0')),float64(len(s)))) 
	}

	return sum == num
}

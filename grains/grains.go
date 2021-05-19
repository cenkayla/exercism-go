package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64,error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("out of range")
	}

	return uint64(math.Pow(2,float64(number-1))),nil
}

func Total() uint64 {
 	var value,total uint64 = 1,1
	for i := 0; i < 64; i++ {
		value *= 2
		total += value
	} 
	return total 
}

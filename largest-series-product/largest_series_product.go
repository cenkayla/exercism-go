package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

func LargestSeriesProduct(digit string, span int) (int, error) {

	if span > len(digit) || span < 0 {
		return -1, errors.New("span must be smaller than string length and span must be greater than zero")
	}
	if len(digit) == 0 && span > 0 {
		return -1, errors.New("span must be smaller than string length")
	}
	largestPro := 0
	d := []rune(digit)
	for i := 0; i <= len(digit)-span; i++ {
		localPro := 1
		for j := i; j < i+span; j++ {
			if unicode.IsLetter(d[j]){
				return -1, errors.New("digits input must only contain digits")
			}
			num, _ := strconv.Atoi(string(digit[j]))
			localPro *= num
		}
		if localPro > largestPro {
			largestPro = localPro
		}
	}
	return largestPro, nil
}
 
package luhn

import (
	"strings"
	"unicode"
)

func Valid(cardNumber string) bool {

	cardNumber = strings.Replace(cardNumber, " ", "", -1)
	if len(cardNumber) < 2 {
		return false
	}

	sum := 0
	isEven := len(cardNumber)%2 == 0
	for _, v := range cardNumber {
		if !unicode.IsNumber(v) {
			return false
		}
		// rune değeri int'e çevirmek için
		num := int(v) - '0'

		if isEven {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		isEven = !isEven
		sum += num
	}
	return sum%10 == 0

}

package romannumerals

import (
	"errors"
	"strings"
)

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(arabicNum int) (string, error) {
	if arabicNum <= 0 || arabicNum > 3000 {
		return "", errors.New("the number must be between 0 and 3000")
	}

	ans := strings.Builder{}

	for _, v := range allRomanNumerals {
		for arabicNum >= v.Value {
			ans.WriteString(v.Symbol)
			arabicNum -= v.Value
		}
	}

	return ans.String(), nil
}

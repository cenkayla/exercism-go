package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)



func gridDimensions(count int) (int, int) {
	c := int(math.Ceil(math.Sqrt(float64(count))))
	r := (count + c - 1) / c
	return r, c
}

func Encode(input string) string {
	re := regexp.MustCompile(`\W+`)
	normalized := re.ReplaceAllString(strings.ToLower(input), "")
	if len(normalized) == 0 {
		return ""
	}

	r, c := gridDimensions(len(normalized))

	var sb strings.Builder
	for row := 0; row < c; row++ {
		for col := 0; col < r; col++ {
			index := (col * c) + row

			if index > len(normalized)-1 {
				sb.WriteByte(' ')
			} else {
				sb.WriteByte(normalized[index])
			}
		}
		if row < c-1 {
			sb.WriteByte(' ')
		}
	}

	return sb.String()
}

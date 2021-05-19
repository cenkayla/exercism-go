package isogram

import "strings"

func IsIsogram(s string) bool {

	input := strings.ToLower(s)
	input = strings.Replace(input, "-", "", -1)
	input = strings.Replace(input, " ", "", -1)

	for i := 0; i < len(input); i++ {
		for y := 1; y < len(input); y++ {
			if input[i] == input[y] {
				if i == y {
					break
				}
				return false
			}
		}
	}
	return true
}
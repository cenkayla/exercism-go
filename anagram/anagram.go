package anagram

import (
	"reflect"
	"sort"
	"strings"
)

func Detect(input string, candidates []string) (ans []string) {

	input = strings.ToLower(input)
	inputSlice := strings.Split(input, "")

	sort.Slice(inputSlice, func(i, j int) bool {
		return inputSlice[i] < inputSlice[j]
	})

	sortedInput := strings.Join(inputSlice, "")

	for _,v := range candidates {
		check := strings.ToLower(v)
		if input == check {
			break
		}
		candidatesSlice := strings.Split(check, "")

		sort.Slice(candidatesSlice, func(i, j int) bool {
			return candidatesSlice[i] < candidatesSlice[j]
		})
		sortedCandidates := strings.Join(candidatesSlice, "")

		if reflect.DeepEqual(sortedInput, sortedCandidates) {
			ans = append(ans, v)
		}
	}
	return ans
}
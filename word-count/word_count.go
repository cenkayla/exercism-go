package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

 	result := Frequency{}
	re := regexp.MustCompile(`\w+('\w)*`)
	input := re.FindAllString(strings.ToLower(phrase), -1)
	for _, v := range input {
		result[v]++
	}
	return result
}

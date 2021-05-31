package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(input string) (string, error) {

	re := regexp.MustCompile(`\W+`)
	input = re.ReplaceAllString(input, "")

	if len(input) == 11 && input[0] == '1' {
		input = input[1:]
	}

	if len(input) != 10 {
		return "", errors.New("error")
	}

	if input[0] < '2' || input[3] < '2' {
		return "", errors.New("error")
	}
	return input, nil
}

func AreaCode(n string) (string, error) {
	n, err := Number(n)
	if err != nil {
		return "", errors.New("error")
	}
	return n[:3], nil
}

func Format(n string) (string, error) {
	n, err := Number(n)
	if err != nil {
		return "", errors.New("error")
	}
	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}

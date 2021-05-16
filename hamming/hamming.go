package hamming

import "errors"

func Distance(a, b string) (int, error) {
	flag := 0
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				flag++
			}
		}
		return flag, nil
	}
	return 0, errors.New("error")

}

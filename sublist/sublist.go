package sublist

import "reflect"

type Relation string

const (
	equal     Relation = "equal"
	unequal   Relation = "unequal"
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
)

func Sublist(s, t []int) (r Relation) {

	l, m := len(s), len(t)
	if l*m == 0 {
		switch {
		case l == m:
			return equal
		case l == 0:
			return sublist
		case m == 0:
			return superlist
		}

	}
	n := l - m
	switch {
	case n == 0:
		if reflect.DeepEqual(s, t) {
			return equal
		}
		return unequal
	case n < 0:
		for i := 0; i+l <= m; i++ {
			if reflect.DeepEqual(s, t[i:i+l]) {
				return sublist
			}
		}
		return unequal
	case n > 0:
		for i := 0; i+m <= l; i++ {
			if reflect.DeepEqual(s[i:i+m], t) {
				return superlist
			}
		}
		return unequal
	}
	return 
}

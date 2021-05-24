package flatten

func Flatten(input interface{}) []interface{} {
	flatInterface := []interface{}{}

	switch v := input.(type) {
	case []interface{}:
		for _, d := range v {
			flatInterface = append(flatInterface, Flatten(d)...)
		}
	case int:
		flatInterface = append(flatInterface, v)
	}

	return flatInterface
}

package series

func All(num int, input string) []string {

	ans := []string{}
	for i := 0; i < len(input)-num+1; i++ {
		ans = append(ans, input[i:i+num])
	}
	return ans

}

func UnsafeFirst(num int,input string) string {

	if str, ok := First(num, input); ok {
		return str
	}

	panic(nil)
}


func First(num int, input string) (first string, ok bool) {

	if num == 0 || num > len(input) {
		return first, ok
	}

	return input[:num], true
}
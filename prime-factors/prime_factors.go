package prime

func Factors(num int64) []int64 {

	s := []int64{}
	for i := int64(2); num >= i; i++ {
		for num%i == 0 {
			num = num / i
			s = append(s, i)
		}
	}
	return s
}

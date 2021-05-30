package sieve

func Sieve(num int) []int {

	ans := []int{}
	if num < 2 {
		return ans
	}
	ans = append(ans, 2)

 	for a := 3; a <= num; a+=2 {
		for i := 2; i < a; i++ {
			if a%i == 0 {
				break
			} else {
				if a == i+1 {
					ans = append(ans, a)
				}
			}
		}
	}
	return ans
}

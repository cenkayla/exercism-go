package diffsquares

func SquareOfSum(number int) int {
	//Ardışık sayıların toplamının karesinin formülü : n.(n+1) / 2
	sum := number * (number + 1) / 2
	return sum * sum
}

func SumOfSquares(number int) int {
	//Ardışık tam kare sayıların toplamının formülü : n.(n+1).(2n+1)/6
	sum := number * (number + 1) * (number*2 + 1) / 6
	return sum
}

func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}

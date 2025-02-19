package Simulation

func CountBalls(lowLimit int, highLimit int) int {
	box := make([]int, 46)
	for i := lowLimit; i <= highLimit; i++ {
		bitSum := getBitSum(i)
		box[bitSum]++
	}

	var res int
	for _, v := range box {
		if v > res {
			res = v
		}
	}
	return res
}

func getBitSum(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

//package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	max32 := math.MaxInt32
	min32 := math.MinInt32
	if divisor == -1 {
		if dividend > min32 {
			return -dividend
		} else {
			return max32
		}
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	ans := div(int64(dividend), int64(divisor))
	if sign == 1 {
		if ans > int64(max32) {
			return max32
		}
		return int(ans)
	} else {
		return -int(ans)
	}
}

func div(a, b int64) int64 {
	if a < b {
		return 0
	}
	count := int64(1)
	tmp := b
	for (tmp << 1) <= a {
		count += count
		tmp = tmp << 1
	}
	return count + div(a-tmp, b)
}
func main() {
	dividend := 2147483647
	divisor := 2
	println(divide(dividend, divisor))
}

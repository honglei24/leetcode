// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func reverseInt(x int) int {
	rs := 0
	for x > 0 {
		rs = rs*10 + x%10
		x /= 10
	}
	return rs
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := reverseInt(x)
	//s := strconv.Itoa(x)
	//s1 := reverseString(s)
	return x == y
}

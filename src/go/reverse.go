package src

import "math"

func reverse(x int) int {
	max := math.MaxInt32
	min := math.MinInt32
	ret := int64(0)
	for ; x != 0; x = x / 10 {
		ret = ret*10 + int64(x%10)
	}
	if ret > int64(max) || ret < int64(min) {
		return 0
	} else {
		return int(ret)
	}
}


func reverse_1(x int) int {
	var ret int
	for x != 0 {
		ret = ret*10 + x%10
		if ret > (1<<31-1) || ret < -(1<<31) {
			return 0
		}
		x = x / 10
	}
	return ret
}

func reverse_2(x int) int {
	s := strconv.Itoa(x)
	flag := 0
	if strings.HasPrefix(s, "-") {
		flag = 1
	}
	b := []byte(s)
	l := len(b)
	if flag == 1 {
		for i := 1; i <= l/2; i++ {
			b[i], b[l-i] = b[l-i], b[i]
		}
	} else {
		for i := 0; i < l/2; i++ {
			b[i], b[l-i-1] = b[l-i-1], b[i]
		}
	}
	int64, err := strconv.ParseInt(string(b), 10, 64)

	if int64 > math.MaxInt32 || int64 < math.MinInt32 || err != nil {
		return 0
	}

	return int(int64)
}
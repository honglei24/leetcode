func myAtoi(str string) int {
	if len(str) == 0 || str == "" {
		return 0
	}
	var res int64
	i := 0
	var negative bool
	for ; i < len(str) && str[i] == ' '; i++ {
	}

	if i == len(str) {
		return 0
	}
	if str[i] == '-' {
		i++
		negative = true
	} else if str[i] == '+' {
		i++
	}
	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int64(str[i]-'0')
			if res > math.MaxInt32 && !negative {
				return math.MaxInt32
			} else if res > math.MaxInt32 && negative {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	if negative {
		return -int(res)
	}
	return int(res)
}
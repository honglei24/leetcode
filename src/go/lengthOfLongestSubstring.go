package src

func lengthOfLongestSubstring(s string) int {
	max := 0
	tmp := 0
	m := make(map[int32]int)
	for i, v := range s {
		if value, ok := m[v]; ok {
			if tmp > max {
				max = tmp
			}
			for j := i - tmp; j <= value; j++ {
				delete(m, int32(s[j]))
			}
			tmp = i - value

		} else {
			tmp = tmp + 1
		}
		m[v] = i
	}
	if tmp > max {
		max = tmp
	}
	return max
}

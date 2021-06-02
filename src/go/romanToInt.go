package src

func romanToInt(s string) int {
	ans := 0
	m := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}
	i := 0
	for i < len(s)-1 {
		if v, ok := m[s[i:i+2]]; ok {
			ans += v
			i += 2
		} else if v, ok := m[s[i:i+1]]; ok {
			ans += v
			i += 1
		}
	}
	if i == len(s)-1 {
		if v, ok := m[s[i:i+1]]; ok {
			ans += v
			i += 1
		}
	}
	return ans
}
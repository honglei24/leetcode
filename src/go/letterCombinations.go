func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ans := []string{""}
	var m = map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
	for i := 0; i < len(digits); i++ {
		digit := digits[i : i+1]
		length := len(ans)
		if value, ok := m[digit]; ok {
			for j := 0; j < len(value); j++ {
				for k := 0; k < length; k++ {
					ans = append(ans, ans[k]+string(value[j]))
				}
			}
			ans = ans[length:]
		}
	}

	return ans
}
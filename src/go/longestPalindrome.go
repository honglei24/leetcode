package src

func longestPalindrome(s string) string {
	if s == "" || len(s) == 0 {
		return ""
	}
	ranger := []int{0, 0}
	for i := 0; i < len(s); {
		i = indexOf(s, i, ranger)
	}
	return s[ranger[0]:ranger[1]]
}

func indexOf(s string, start int, ranger []int) int {
	end := start + 1
	for start >= 0 && end < len(s) && s[start] == s[end] {
		end++
	}
	result := end

	for start > 0 && end < len(s) && s[start-1] == s[end] {
		start--
		end++
	}

	if end-start > ranger[1]-ranger[0] {
		ranger[0] = start
		ranger[1] = end
	}
	return result
}
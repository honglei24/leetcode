package src

func generateParenthesis(n int) []string {
	var ans []string
	var s []byte
	backtrack(&ans, s, 0, 0, n)
	return ans
}

func backtrack(ans *[]string, s []byte, l int, r int, max int) {
	if len(s) == 2*max {
		*ans = append(*ans, string(s))
		return
	}
	if l < max {
		s = append(s, '(')
		backtrack(ans, s, l+1, r, max)
		s = s[:len(s)-1]
	}
	if r < l {
		s = append(s, ')')
		backtrack(ans, s, l, r+1, max)
		s = s[:len(s)-1]
	}
}
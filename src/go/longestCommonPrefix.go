func longestCommonPrefix(strs []string) string {
	if nil == strs || len(strs) == 0 {
		return ""
	}
	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(ans) && j < len(strs[i]); j++ {
			if ans[j:j+1] != strs[i][j:j+1] {
				break
			}
		}
		ans = ans[0:j]
		if ans == "" {
			return ""
		}
	}
	return ans
}
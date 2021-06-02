package src

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	isFirstMatch := false
	if s != "" && (s[0] == p[0] || strings.HasPrefix(p, ".")) {
		isFirstMatch = true
	}
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (isFirstMatch && isMatch(s[1:], p))
	} else {
		return isFirstMatch && isMatch(s[1:], p[1:])
	}
	return true
}
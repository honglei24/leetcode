func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}

	m := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	var stack []string
	for i := 0; i < length; i++ {
		tmp := s[i : i+1]
		if tmp == " " {
			continue
		}
		if _, ok := m[tmp]; ok {
			stack = append(stack, tmp)
		} else {
			lStack := len(stack)
			if len(stack) > 0 {
				if v, ok := m[stack[lStack-1]]; ok {
					if v == s[i:i+1] {
						stack = stack[:lStack-1]
					} else {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
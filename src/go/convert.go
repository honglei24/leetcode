package src

func convert(s string, numRows int) string {
	ret := ""
	if s == "" || numRows < 1 {
		return ret
	}
	if numRows == 1 {
		return s
	}
	var tmp []string = make([]string, numRows)
	for i := 0; i < len(s); i++ {
		r := i / (numRows - 1)
		c := i % (numRows - 1)
		if r%2 == 0 {
			tmp[c] = tmp[c] + s[i:i+1]
		} else {
			tmp[numRows-c-1] = tmp[numRows-c-1] + s[i:i+1]
		}
	}
	for j := 0; j < numRows; j++ {
		ret = ret + tmp[j]
	}
	return ret
}
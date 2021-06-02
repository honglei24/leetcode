//package main

func findSubstring(s string, words []string) []int {
	ans := []int{}
	lenS := len(s)
	wordNum := len(words)
	if s == "" || lenS == 0 || words == nil || wordNum == 0 {
		return ans
	}
	oneWord := len(words[0])
	allLen := oneWord * wordNum
	if lenS < allLen {
		return ans
	}

	m := make(map[string]int)

	for _, word := range words {
		m[word] = getOrDefault(m, word) + 1
	}
	for i := 0; i < oneWord; i++ {
		left := i
		right := i
		count := 0

		tmp := make(map[string]int)
		for right+oneWord <= lenS {
			w := s[right : right+oneWord]
			right += oneWord
			if _, ok := m[w]; !ok {
				count = 0
				left = right
				tmp = make(map[string]int)
			} else {
				tmp[w] = getOrDefault(tmp, w) + 1
				count++
				for getOrDefault(tmp, w) > getOrDefault(m, w) {
					leftWord := s[left : left+oneWord]
					tmp[leftWord] = getOrDefault(tmp, leftWord) - 1
					count--
					left += oneWord
				}
			}
			if count == wordNum {
				ans = append(ans, left)
			}
		}
	}
	return ans
}

func getOrDefault(m map[string]int, key string) int {
	if v, ok := m[key]; ok {
		return v
	} else {
		return 0
	}
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	ans := findSubstring(s, words)
	for _, v := range ans {
		println(v)
	}
}

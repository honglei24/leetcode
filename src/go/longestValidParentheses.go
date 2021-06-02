package src

//package main

func longestValidParentheses(s string) int {
	ans := 0
	length := len(s)
	d := make([]int, length)
	for i := 1; i < length; i++ {
		if s[i:i+1] == ")" {
			if s[i-1:i] == "(" {
				if i >= 2 {
					d[i] = 2 + d[i-2]
				} else {
					d[i] = 2
				}
			} else if i-d[i-1]-1 >= 0 && s[i-d[i-1]-1:i-d[i-1]] == "(" {
				if i-d[i-1] >= 2 {
					d[i] = 2 + d[i-1] + d[i-d[i-1]-2]
				} else {
					d[i] = 2 + d[i-1]
				}
			}
		}
		if d[i] > ans {
			ans = d[i]
		}
	}
	return ans
}

func main() {
	s := ")()())"
	println(longestValidParentheses(s))
}

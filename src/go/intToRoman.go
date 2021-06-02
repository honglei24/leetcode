package src

func intToRoman(num int) string {
	ret := ""
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(nums); i++ {
		for num >= nums[i] {
			ret = ret + romans[i]
			num = num - nums[i]
		}
	}
	return ret
}

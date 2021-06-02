package src

func threeSumClosest(nums []int, target int) int {
	if nums == nil || len(nums) <= 2 {
		return 0
	}
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	if ans >= target {
		return ans
	}
	for k := 0; k < len(nums)-2; k++ {
		i := k + 1
		j := len(nums) - 1

		for i < j {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			} else if sum < target {
				if sum > ans || target-sum < ans-target {
					ans = sum
				}
				i++
			} else {
				if sum-target < target-ans || sum < ans {
					ans = sum
				}
				j--
			}

		}
	}
	return ans
}
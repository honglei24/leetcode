package src

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	if nums == nil || len(nums) <= 2 {
		return nil
	}
	sort.Ints(nums)
	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i := k + 1
		j := len(nums) - 1

		for i < j {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				i++
				for nums[i] == nums[i-1] && i < j {
					i++
				}
				j--
				for nums[j] == nums[j+1] && i < j {
					j--
				}
			} else if sum < 0 {
				i++
				for nums[i] == nums[i-1] && i < j {
					i++
				}
			} else {
				j--
				for nums[j] == nums[j+1] && i < j {
					j--
				}
			}

		}
	}
	return ans
}
func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	if nums == nil || len(nums) < 4 {
		return ans
	}
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		minSum := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
		if minSum > target {
			break
		}
		maxSum := nums[i] + nums[length-3] + nums[length-2] + nums[length-1]
		if maxSum < target {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			h := length - 1

			minSum = nums[i] + nums[j] + nums[l] + nums[l+1]
			if minSum > target {
				break
			}
			maxSum = nums[i] + nums[j] + nums[h-1] + nums[h]
			if maxSum < target {
				continue
			}
			for l < h {
				curr := nums[i] + nums[j] + nums[l] + nums[h]
				if curr == target {
					tmp := []int{nums[i], nums[j], nums[l], nums[h]}
					ans = append(ans, tmp)
					l++
					for l < h && nums[l] == nums[l-1] {
						l++
					}
					h--
					for l < h && nums[h] == nums[h+1] {
						h--
					}
				} else if curr < target {
					l++
				} else {
					h--
				}
			}
		}
	}
	return ans
}
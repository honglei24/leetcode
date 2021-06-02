package src

//package main

func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := length - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	left := i + 1
	right := length - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	for _, v := range nums {
		println(v)
	}
}

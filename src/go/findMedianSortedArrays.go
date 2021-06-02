import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	left := (l1 + l2 + 1) / 2
	right := (l1 + l2 + 2) / 2
	return float64(findk(nums1, 0, nums2, 0, left)+findk(nums1, 0, nums2, 0, right)) / 2.0
}

func findk(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return int(math.Min(float64(nums1[i]), float64(nums2[j])))

	}
	mid1, mid2 := math.MaxInt64, math.MaxInt64
	if i+k/2-1 < len(nums1) {
		mid1 = nums1[i+k/2-1]
	}
	if j+k/2-1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	}
	if mid1 > mid2 {
		return findk(nums1, i, nums2, j+k/2, k-k/2)
	} else {
		return findk(nums1, i+k/2, nums2, j, k-k/2)
	}
}
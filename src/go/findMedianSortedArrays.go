func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	nums3 := make([]int, l)
	index1, index2 := 0, 0
	for k := 0; k < l; k++ {
		if index1 < len(nums1) && index2 < len(nums2) {
			if nums1[index1] > nums2[index2] {
				nums3[k] = nums2[index2]
				index2++
			} else {
				nums3[k] = nums1[index1]
				index1++
			}
		} else if index1 >= len(nums1) {
			nums3[k] = nums2[index2]
			index2++
		} else {
			nums3[k] = nums1[index1]
			index1++
		}
	}
	if l%2 == 0 {
		return float64(nums3[l/2-1]+nums3[l/2]) / 2.0
	} else {
		return float64(nums3[(l-1)/2])
	}
}
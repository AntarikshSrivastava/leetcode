func nextGreaterElement(nums1 []int, nums2 []int) []int {

	for i, n := range nums1 {
		val := search(n, nums2)
		if val != -1 {
			found := false
			for j := val + 1; j < len(nums2); j++ {
				if nums2[j] > n {
					found = true
					nums1[i] = nums2[j]
					break
				}
			}
			if !found {
				nums1[i] = -1
			}
		}
	}
	return nums1
}

func search(n int, nums []int) int {
	for k, v := range nums {
		if v == n {
			return k
		}
	}
	return -1
}
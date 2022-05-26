func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := map[int]bool{}
	m2 := map[int]bool{}

	for _, n1 := range nums1 {
		m1[n1] = false
	}
	for _, n2 := range nums2 {
		m2[n2] = false
	}

	res1 := []int{}
	res2 := []int{}

	for _, n1 := range nums1 {

		if _, ok := m2[n1]; !ok {
			res1 = append(res1, n1)
			m2[n1] = true
		}
	}

	for _, n2 := range nums2 {
		if _, ok := m1[n2]; !ok {
			res2 = append(res2, n2)
			m1[n2] = true
		}
	}

	res := [][]int{}
	res = append(res, res1)
	res = append(res, res2)

	return res
}
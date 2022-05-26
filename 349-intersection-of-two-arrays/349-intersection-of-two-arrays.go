func contains(arr []int, t int) bool {
	for _, a := range arr {
		if a == t {
			return true
		}
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	flags := map[int]bool{}
	for _, n := range nums1 {
		flags[n] = true
	}

	var intersection []int
	for _, m := range nums2 {
		if flags[m] && !contains(intersection, m) {
			intersection = append(intersection, m)
		}
	}
	return intersection
}
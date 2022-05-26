func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 0 {
		return res
	}

	start := nums[0]
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(start))
		return res
	}
	end := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if start == end {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			start = nums[i]
		}
		end = nums[i]
	}

	if start == end {
		res = append(res, strconv.Itoa(start))
	} else {
		res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
	return res
}
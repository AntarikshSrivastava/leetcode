func findRelativeRanks(score []int) []string {
	m := make(map[int]int, len(score))
	sc := []int{}

	for _, v := range score {
		sc = append(sc, v)
	}

	sort.Slice(score, func(i, j int) bool {
		return score[j] < score[i]
	})

	for i, v := range score {
		m[v] = i
	}
	
	res := []string{}
	for _, v := range sc {
		switch m[v] {
		case 0:
			res = append(res, "Gold Medal")
		case 1:
			res = append(res, "Silver Medal")
		case 2:
			res = append(res, "Bronze Medal")
		default:
			res = append(res, strconv.Itoa(m[v]+1))
		}
	}
	return res
}
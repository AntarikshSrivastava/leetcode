var m map[int]int

func findMode(root *TreeNode) []int {
	m = map[int]int{}
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	countFreq(root)
	var res []int
	max := math.MinInt32
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	for k, v := range m {
		if v == max {
			res = append(res, k)
		}
	}
	return res
}

func countFreq(root *TreeNode) {
	if root == nil {
		return
	}
	if _, ok := m[root.Val]; !ok {
		m[root.Val] = 1
	} else {
		m[root.Val]++
	}
	countFreq(root.Left)
	countFreq(root.Right)
}
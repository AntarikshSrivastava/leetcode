func matrixReshape(mat [][]int, r int, c int) [][]int {
	queue := []int{}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			queue = append(queue, mat[i][j])
		}
	}

	l := len(queue)

	if r*c != l {
		return mat
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
        res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			first := queue[0]
			queue = queue[1:]
			res[i][j] = first
		}
	}

	return res
}
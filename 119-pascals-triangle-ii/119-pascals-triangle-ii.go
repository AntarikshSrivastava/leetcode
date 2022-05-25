func getRow(rowIndex int) []int {
    numRows := rowIndex + 1
	res := make([][]int, numRows)
    
	for i := 1; i <= numRows; {
		j := i
		for j <= numRows {
			res[j-1] = append(res[j-1], 1)
			j++
		}
		i++
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
    fmt.Println(res)
    return res[numRows-1]
}
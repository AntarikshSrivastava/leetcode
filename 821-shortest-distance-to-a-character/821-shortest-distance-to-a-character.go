func shortestToChar(s string, c byte) []int {
    
    idx := []int{}
    res := make([]int, len(s))
    
    for i, v := range []byte(s) {
        if v == c {
            idx = append(idx, i)
        }
    }
    fmt.Println("idx: ", idx)
    
    for i, v := range []byte(s) {
        if v != c {
            max := 10001
            for _, val := range idx {
                diff := diffAbs(i, val)
                if max > diff {
                    max = diff
                }
            }
            res[i] = max
        } else {
            res[i] = 0
        }
    }
    return res
}

func diffAbs(a, b int) int {
    if a - b < 0 {
        return b - a
    }
    return a - b
}
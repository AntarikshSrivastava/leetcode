func numJewelsInStones(jewels string, stones string) int {
    jm := make(map[rune]bool)
    res := 0
    for _, j := range jewels {
        jm[j] = true
    }
    
    for _, s := range stones {
        if _, ok := jm[s]; ok {
            res++
        }
    }
    return res
}
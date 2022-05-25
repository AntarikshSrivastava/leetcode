func timeRequiredToBuy(tickets []int, k int) int {
    val := tickets[k]
    res := 0
    
    for i := 0; i < len(tickets); i++ {
        res += min(tickets[i], val)
        
        if i == k {
            val--
        }
    }
    return res
}

func min(a, b int) int {
    if a > b { 
        return b 
    }
    return a
}
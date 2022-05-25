func maxProfit(prices []int) int {
    
    if len(prices) == 1 {
        return 0
    }
    
	maxSoFar := 0
    maxEndingHere := 0
    
    for i:=1; i<len(prices); i++ {
        maxEndingHere = maxEndingHere + prices[i] - prices[i-1]
        
        if maxSoFar < maxEndingHere {
            maxSoFar = maxEndingHere
        }
        
        if maxEndingHere < 0 {
            maxEndingHere = 0
        }
    }

	return maxSoFar
}

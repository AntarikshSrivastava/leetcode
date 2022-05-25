func finalPrices(prices []int) []int {
	stack := []int{}
	for i, p := range prices {

		for len(stack) != 0 && prices[stack[len(stack)-1]] >= p {
			prices[stack[len(stack)-1]] -= p
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return prices
}
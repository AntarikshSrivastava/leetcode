func maxDepth(s string) int {
	stack := []rune{}
	depth := 0
	maxD := 0
	for _, c := range s {
		if c == ')' || c == '(' {
			n := len(stack)
			if n > 0 && ((c == ')' && stack[n-1] == '(') || (c == ')' && stack[n-1] == '(')) {
				// stack = stack[:n-1]
				depth--

			} else {
				stack = append(stack, c)
				depth++
				if depth >= maxD {
					maxD = depth
				}
			}
		}
	}
	return maxD
}
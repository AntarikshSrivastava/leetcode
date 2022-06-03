func isPowerOfTwo(n int) bool {
    if n <= 0{
        return false
    }
	if n == 1 || n== 2{
		return true
	}
	for n > 2 {
		if n%2 == 0 {
			n /= 2
		} else {
			return false
		}
	}
	return true

}
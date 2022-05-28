func sortedSquares(nums []int) []int {
    l := len(nums)
    left := 0
    right := l - 1
    res := make([]int, l)
    
    for i := l-1; i>=0; i-- {
        if abs(nums[left]) < abs(nums[right]) {
            res[i] = nums[right]*nums[right]
            right--
        } else {
            res[i] = nums[left]*nums[left]
            left++
        }
    }
    return res
}



func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
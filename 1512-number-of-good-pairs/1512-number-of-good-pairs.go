func numIdenticalPairs(nums []int) int {
    ans, tmp := 0, make([]int, 101)
    for _, n := range nums { tmp[n]++ }
    for _, v := range tmp { if v > 1 { ans += v * (v - 1) / 2 } }
    return ans
}
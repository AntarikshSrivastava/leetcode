func findDisappearedNumbers(nums []int) []int {
    res := []int{}
    
    for i:=0; i<len(nums); i++ {
        j:= abs(nums[i])-1
        nums[j] = abs(nums[j]) * -1
    }
    
    for k, v := range nums {
        if v > 0 {
            res = append(res, k+1)
        }
    }
    return res
}

func abs(n int) int {
    if n > 0 { return n }
    return -n
}
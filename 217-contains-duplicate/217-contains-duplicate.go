func containsDuplicate(nums []int) bool {
    m := map[int]int{}   
    
    for _, val := range nums {
        m[val]++
        if m[val] > 1 {
            return true
        }
    }
    
    return false
}
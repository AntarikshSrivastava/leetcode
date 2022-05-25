func containsDuplicate(nums []int) bool {
    sort.Ints(nums) // Need O(N*log(N))
    if len(nums)<2{return false}
    for i:=1;i<len(nums);i++{
        if nums[i]==nums[i-1] {return true}        
    }
    return false
}
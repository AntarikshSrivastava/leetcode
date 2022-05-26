func intersection(nums [][]int) []int {
    
    l := len(nums)
    if l == 1 {
        res := nums[0]
        sort.Ints(res)
        return res
    } 
    
    m := map[int]int{}
    for _, v := range nums[0] {
        m[v] = 1    
    }
    
    for i:=1; i<l; i++ {
        for _, n := range nums[i] {
            if _, ok := m[n]; ok {
                m[n]++
            }
        }
    }
    
    res := []int{}
    
    for k, v := range m {
        if v == l {
            res = append(res, k)
        }
    }
    sort.Ints(res)
    return res
}
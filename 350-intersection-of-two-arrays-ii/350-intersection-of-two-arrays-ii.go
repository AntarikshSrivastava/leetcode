func intersect(nums1 []int, nums2 []int) []int { 
    var intersection []int
    m := map[int]int{}
    
    for _, n := range nums1 {
        m[n]++
    }
    
    for _, v := range nums2 {
        if m[v] > 0 {
            intersection = append(intersection, v)
            m[v]--
        }
    }
    return intersection
}
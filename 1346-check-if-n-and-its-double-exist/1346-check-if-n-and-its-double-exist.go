func checkIfExist(arr []int) bool {
    m := make(map[int]bool)
    
    for _, n := range arr {
        if n%2==0 {
            if _, ok := m[2*n]; ok {
                return true
            }
            
            if _, ok := m[n/2]; ok {
                return true
            }
        } else {
            if _, ok := m[2*n]; ok {
                return true
            }
        }
        m[n] = true
    }
    return false
}
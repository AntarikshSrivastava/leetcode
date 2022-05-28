func duplicateZeros(arr []int)  {
    
    stack := make([]int, 0)
    i := 0
    for len(stack) != len(arr) {
        if arr[i] != 0 {
            stack = append(stack, arr[i])
        } else {
            stack = append(stack, 0)
            stack = append(stack, 0)
            
            if len(stack) > len(arr) {
                stack = stack[:len(arr)]
            }
        }
        i++
    }
    
    for i:=0; i<len(arr); i++ {
        arr[i] = stack[i]
    }
}
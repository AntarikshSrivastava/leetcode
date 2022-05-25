func minOperations(logs []string) int {
    stack := []string{}
    for _, log := range logs {
        if log == "../" {
            if len(stack) > 0 {
                stack = stack[1:]
            }    
        } else if log != "./" {
            stack = append(stack, log)    
        }
    }
    return len(stack)
}
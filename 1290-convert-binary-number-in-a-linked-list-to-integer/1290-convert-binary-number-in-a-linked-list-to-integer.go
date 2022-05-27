/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    curr := head
    
    n := 0
    for curr != nil {
        curr = curr.Next
        n++
    }
    
    curr = head
    sum := 0
    
    for curr != nil {
        sum += curr.Val * int(math.Pow(2, float64(n-1)))
        n--
        curr = curr.Next
    }
    return sum
}
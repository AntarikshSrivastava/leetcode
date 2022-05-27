/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    ret := 0
    for ; head != nil; head = head.Next {
        fmt.Println(ret, ret << 1, head.Val, ret << 1 + head.Val)
        ret = ret << 1 + head.Val
    }
    return ret
}
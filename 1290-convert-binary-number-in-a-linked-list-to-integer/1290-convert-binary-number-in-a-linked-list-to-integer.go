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
        ret = ret << 1 + head.Val
    }
    return ret
}
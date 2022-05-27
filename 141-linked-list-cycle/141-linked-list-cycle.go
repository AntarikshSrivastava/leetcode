/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := head.Next
    
    for fast != slow && fast != nil && fast.Next != nil {
        fastNext := fast.Next.Next
        slowNext := slow.Next
        fast = fastNext
        slow = slowNext
    }
    return fast == slow
}
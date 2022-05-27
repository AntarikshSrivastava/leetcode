/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    slower := head
    faster := head
    for faster != nil && faster.Next != nil {
        slower = slower.Next
        faster = faster.Next.Next
        if slower == faster {
            return true
        }
    } 
    return false
}
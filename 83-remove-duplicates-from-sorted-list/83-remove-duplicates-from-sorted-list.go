/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    curr := head
    for curr != nil {
        next := curr.Next
        for next != nil && next.Val==curr.Val {
            next = next.Next
        }
        curr.Next = next
        curr = next
    }
    
    return head
}
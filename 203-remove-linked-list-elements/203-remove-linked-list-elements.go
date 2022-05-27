/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }
    
    for head != nil && head.Val == val {
		head = head.Next
	}
    
    curr := head 
    prev := head
    
    for curr != nil {
        if val == curr.Val {
            prev.Next = curr.Next
        } else {
            prev = curr
        }
        curr = curr.Next
    }
    return head
}
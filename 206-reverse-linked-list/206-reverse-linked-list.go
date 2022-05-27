/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode  {
	var prev *ListNode
	var curr *ListNode = head
	for curr != nil {
        fmt.Println("xxxx: ", curr)
        next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next 
        
        fmt.Println("head: ", head, "curr: ", curr, "prev: ", prev, "next: ", next) 
	}
	return prev
}
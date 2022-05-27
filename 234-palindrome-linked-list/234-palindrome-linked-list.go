/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    
    if head.Next == nil {
        return true
    }
    
    curr := head
    
    n := 0
    for curr != nil {
        curr = curr.Next
        n++
    }
    
    curr = head
    var stop *ListNode
    
    firstHalf:=0
    for firstHalf != n/2 {
        stop = curr
        curr=curr.Next   
        firstHalf++
    }
    

    
    if n%2 != 0 {
        stop = curr
        curr = curr.Next 
    } 
    
    var prev *ListNode
    
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
        
        if curr == nil {
            stop.Next = prev
        }
    }
    
    curr = head
    stop = stop.Next
    
    for curr!=nil && stop != nil {
        if curr.Val != stop.Val {
            return false
        }
        curr = curr.Next
        stop = stop.Next
    }
    
    return true
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
    var prev *ListNode
    
    for list1 != nil && list2 != nil {
        
        var curr *ListNode
        
        if list1.Val < list2.Val {
            curr = list1
            list1 = list1.Next
        } else {
            curr = list2
            list2 = list2.Next
        }
        
        if prev == nil {
            head = curr    
        } else {
            prev.Next = curr
        }
        prev = curr
    }
    
    for list1 != nil {
        if prev == nil {
            head = list1
        } else {
            prev.Next = list1
        }
        prev = list1
        list1 = list1.Next
    }
    
    for list2 != nil {
        if prev == nil {
            head = list2
        } else {
            prev.Next = list2
        }
        prev = list2
        list2 = list2.Next
    }
    
    return head
}
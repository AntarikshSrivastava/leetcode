/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]bool) //build a map to cache the nodes
    cur := head
    for cur!=nil{
        if m[cur]==true {return true} // if cur is in the map,there is a circle.
        m[cur]=true
        cur=cur.Next
    }
    return false
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    arr := []*TreeNode{}
    inorder(root, &arr)
    
    res := arr[0]
    curr := res
    
    for _, n := range arr[1:] {
        curr.Left = nil
        curr.Right = n
        curr = n
    }
    curr.Right = nil
    curr.Left = nil
    return res
}

func inorder(root *TreeNode, arr *[]*TreeNode) {
    if root == nil {
        return
    }
    inorder(root.Left, arr)
    *arr = append(*arr, root)
    inorder(root.Right, arr)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    return helper(root, root.Val)
}

func helper(root *TreeNode, val int) bool {
    if root == nil {
        return true
    }
    
    if root.Val != val {
        return false
    }
    
    return helper(root.Left, root.Val) && helper(root.Right, root.Val)
}
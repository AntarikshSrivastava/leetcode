/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
        return true
    }
    if root == nil {
        return false
    }
    return helper(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func helper(root *TreeNode, subRoot *TreeNode) bool{
    if root == nil && subRoot == nil {
        return true
    }
    
    if root == nil || subRoot == nil {
        return false
    }
    
    return root.Val == subRoot.Val && helper(root.Left, subRoot.Left) && helper(root.Right, subRoot.Right)   
}
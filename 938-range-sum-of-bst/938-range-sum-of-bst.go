/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    res := 0
    helper(root, low, high, &res)
    return res
}

func helper(root *TreeNode, low, high int, res *int) {
    if root == nil {
        return
    }
    
    if low <= root.Val && root.Val <= high {
        *res += root.Val
        helper(root.Left, low, high, res)
        helper(root.Right, low, high, res)
        return
    }
    
    if root.Val < low {
        helper(root.Right, low, high, res)
    } else {
        helper(root.Left, low, high, res)
    }
}
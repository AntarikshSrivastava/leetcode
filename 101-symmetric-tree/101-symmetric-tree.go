/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
    if left == nil || right == nil {
        return left == right
    }
    return left.Val == right.Val && helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
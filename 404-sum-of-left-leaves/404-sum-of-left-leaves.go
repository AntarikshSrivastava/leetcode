/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0
    helper(root, &sum, false)
    return sum
}

func helper(root *TreeNode, sum *int, isLeft bool) {
    if root == nil {
        return
    }
    
    if root.Left == nil && root.Right == nil && isLeft {
        *sum += root.Val
    }
    
    helper(root.Left, sum, true)
    helper(root.Right, sum, false)
}
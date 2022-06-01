/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var prev *TreeNode

func inOrder(node *TreeNode) {

    if node == nil {
        return
    }
    
    inOrder(node.Left)
    
    prev.Right = node
    prev = node
    node.Left = nil
    
    inOrder(node.Right)
    
}

func increasingBST(root *TreeNode) *TreeNode {
    
    ans := &TreeNode{-1, nil, nil}
    prev = ans
    
    inOrder(root)
    
    return ans.Right
}
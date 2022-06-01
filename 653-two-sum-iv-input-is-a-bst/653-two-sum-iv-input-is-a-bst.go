/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    nums := make([]int, 0)
    inorder(root, &nums)
    
    for i, j := 0, len(nums)-1; i<j; {
        val := nums[i] + nums[j]
        if val == k {
            return true
        } else if val > k {
            j--
        } else {
            i++
        }
    }
    return false
}

func inorder(root *TreeNode, nums *[]int) {
    if root == nil {
        return
    }
    
    inorder(root.Left, nums)
    *nums = append(*nums, root.Val)
    inorder(root.Right, nums)
}
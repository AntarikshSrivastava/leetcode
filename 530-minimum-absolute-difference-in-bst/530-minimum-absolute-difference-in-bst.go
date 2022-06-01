/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    nums := make([]int, 0)
    inorder(root, &nums)
    min := nums[1] - nums[0]
    
    for i:=2; i<len(nums); i++ {
        if nums[i] - nums[i-1] < min {
            min = nums[i] - nums[i-1]
        }
    }
    return min
}

func inorder(root *TreeNode, nums *[]int) {
    if root == nil {
        return
    }
    inorder(root.Left, nums)
    *nums = append(*nums, root.Val)
    inorder(root.Right, nums)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	val1, val2 := p.Val, q.Val
	cur := root
	for {
		switch {
		case val1 < cur.Val && val2 < cur.Val:
			cur = cur.Left
		case val1 > cur.Val && val2 > cur.Val:
			cur = cur.Right
		default:
			return cur
		}
	}
	return root
}
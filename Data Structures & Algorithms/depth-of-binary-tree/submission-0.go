/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left) + 1
	rightMax := maxDepth(root.Right) + 1
	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
	
}
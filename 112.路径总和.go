/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	target := sum - root.Val
	if root.Left == nil && root.Right == nil {
		if target == 0 {
			return true
		} else {
			return false
		}
	}

	var l, r bool
	if root.Left != nil {
		l = hasPathSum(root.Left, target)
	}
	if root.Right != nil {
		r = hasPathSum(root.Right, target)
	}

	return l || r
}

// @lc code=end


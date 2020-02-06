/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end


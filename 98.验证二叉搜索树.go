/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	return checkBST(root, nil, nil)
}

func checkBST(root *TreeNode, low *int, high *int) bool {
	if root == nil {
		return true
	}

	if low != nil && root.Val <= *low {
		return false
	}

	if high != nil && root.Val >= *high {
		return false
	}

	v := root.Val
	return checkBST(root.Left, low, &v) && checkBST(root.Right, &v, high)
}

// @lc code=end


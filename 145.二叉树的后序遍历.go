/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	type ComNode struct {
		Color int8
		Node  *TreeNode
	}
	const WHITE = 1
	const GRAY = 0

	ret := []int{}
	stack := []*ComNode{&ComNode{Color: WHITE, Node: root}}
	for len(stack) != 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[0:n]

		if node.Color == WHITE {
			stack = append(stack, &ComNode{Color: GRAY, Node: node.Node})
			if node.Node.Right != nil {
				stack = append(stack, &ComNode{Color: WHITE, Node: node.Node.Right})
			}
			if node.Node.Left != nil {
				stack = append(stack, &ComNode{Color: WHITE, Node: node.Node.Left})
			}
		} else {
			ret = append(ret, node.Node.Val)
		}
	}
	return ret
}

// @lc code=end


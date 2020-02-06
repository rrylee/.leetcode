/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	const WHITE = 1
	const GRAY = 0

	type ComNode struct {
		Color uint8
		Node  *TreeNode
	}

	stack := []*ComNode{&ComNode{Color: WHITE, Node: root}}
	for len(stack) != 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[0:n]
		if node.Color == WHITE {
			if node.Node.Right != nil {
				stack = append(stack, &ComNode{Color: WHITE, Node: node.Node.Right})
			}
			stack = append(stack, &ComNode{Color: GRAY, Node: node.Node})
			if node.Node.Left != nil {
				stack = append(stack, &ComNode{Color: WHITE, Node: node.Node.Left})
			}
		} else {
			ret = append(ret, node.Node.Val)
		}
	}

	return ret
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}

// @lc code=end


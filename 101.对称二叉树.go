/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			// Right: &TreeNode{
			// 	Val: 4,
			// },
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
			// Left: &TreeNode{
			// 	Val: 4,
			// },
		},
	}
	fmt.Println(isSymmetric(root))
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	return n1.Val == n2.Val && isMirror(n1.Right, n2.Left) && isMirror(n1.Left, n2.Right)
}

// @lc code=end

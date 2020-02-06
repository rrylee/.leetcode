/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	ret := [][]int{}

	for len(nodes) > 0 {
		tempNodes := []*TreeNode{}
		tempRet := []int{}
		for _, node := range nodes {
			tempRet = append(tempRet, node.Val)

			if node.Left != nil {
				tempNodes = append(tempNodes, node.Left)
			}
			if node.Right != nil {
				tempNodes = append(tempNodes, node.Right)
			}
		}
		ret = append(ret, tempRet)
		nodes = tempNodes
	}

	newRet := [][]int{}
	for i := len(ret) - 1; i >= 0; i-- {
		newRet = append(newRet, ret[i])
	}
	return newRet
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
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
func levelOrder(root *TreeNode) [][]int {
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

	return ret
}

// @lc code=end

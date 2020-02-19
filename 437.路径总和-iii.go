/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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

func pathSum(root *TreeNode, sum int) int {
	return sums(root, sum, make([]int, 1000, 1000), 0)
}

func sums(node *TreeNode, sum int, arr []int, p int) int {
	if node == nil {
		return 0
	}

	n := 0
	if node.Val == sum {
		n += 1
	}
	temp := node.Val
	for i := p - 1; i >= 0; i-- {
		temp += arr[i]
		if temp == sum {
			n += 1
		}
	}

	arr[p] = node.Val
	return n + sums(node.Left, sum, arr, p+1) + sums(node.Right, sum, arr, p+1)
}

//以每个根结点为头，依次计算
func pathSum1(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == sum {
		res += 1
	}

	res += dfs(root.Left, sum-root.Val)
	res += dfs(root.Right, sum-root.Val)

	return res
}

// @lc code=end

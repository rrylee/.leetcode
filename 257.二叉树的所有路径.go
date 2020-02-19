import "fmt"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := []string{}
	dfs(root, "", &res)
	return res
}

func dfs(node *TreeNode, path string, res *[]string) {
	if node.Left == nil && node.Right == nil {
		path += fmt.Sprintf("%d", node.Val)
		*res = append(*res, path)
	} else {
		if node.Left != nil {
			dfs(node.Left, path+fmt.Sprintf("%d->", node.Val), res)
		}
		if node.Right != nil {
			dfs(node.Right, path+fmt.Sprintf("%d->", node.Val), res)
		}
	}
}

// @lc code=end


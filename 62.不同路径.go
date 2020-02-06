/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	dp := [][]int{}
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n]
}

// @lc code=end


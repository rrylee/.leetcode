/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	// dp := make([]int, n+1)
	// dp[0] = 0
	// dp[1] = 1
	// dp[2] = 2
	// for index := 3; index <= n; index++ {
	// 	dp[index] = dp[index-1] + dp[index-2]
	// }
	// return dp[n]
	if n == 2 {
		return 2
	}
	lastOne := 2
	lastTwo := 1
	sum := 0
	for index := 3; index <= n; index++ {
		sum = lastOne + lastTwo
		lastTwo = lastOne
		lastOne = sum
	}
	return sum
}

// @lc code=end


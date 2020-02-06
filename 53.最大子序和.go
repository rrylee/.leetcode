/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	curMax := nums[0]

	for i := 1; i < len(nums); i++ {
		if curMax < 0 {
			curMax = 0
		}
		curMax += nums[i]

		if curMax > max {
			max = curMax
		}
	}

	return max
}

// dp[i] = max(dp[i-1] + num[i], num[i])
// dp[i] = max(dp[i-1], 0) + num[i]

// @lc code=end


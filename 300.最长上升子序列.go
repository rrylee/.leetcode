/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for index, _ := range nums {
		dp[index] = 1
	}

	max := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	exceptedSum := len(nums) * (1 + len(nums)) / 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return exceptedSum - actualSum
}

func missingNumber1(nums []int) int {
	miss := len(nums)

	for i := 0; i < len(nums); i++ {
		miss ^= i ^ nums[i]
	}

	return miss
}

// 0 1 3
// 3 ^ 0 ^ 0 ^ 1 ^ 1 ^ 2 ^ 3

// @lc code=end


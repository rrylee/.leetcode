package main

func main() {
	firstMissingPositive([]int{1, 2, 3})
}

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	num := 1
	for i := 0; i < len(nums); i++ {
		if num == nums[i] {
			num += 1
			i = -1
		}
		if i == len(nums)-1 {
			break
		}
	}
	return num
}

// @lc code=end

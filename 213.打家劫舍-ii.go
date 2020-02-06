/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	nums2 := make([]int, n)
	copy(nums2, nums)

	return maxf(rob2(nums, 0, n-1), rob2(nums2, 1, n-1))
}

func rob2(nums []int, start, end int) int {
	//
}

func maxf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

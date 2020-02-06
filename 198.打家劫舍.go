// package main

// /*
//  * @lc app=leetcode.cn id=198 lang=golang
//  *
//  * [198] 打家劫舍
//  */

// func main() {
// 	rob([]int{2, 1, 1, 2})
// }

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	for i := 1; i < n; i++ {
		if i == 1 {
			nums[i] = maxf(nums[i], nums[i-1])
		} else {
			nums[i] = maxf(nums[i-2]+nums[i], nums[i-1])
		}
	}

	return nums[n-1]
}

func maxf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

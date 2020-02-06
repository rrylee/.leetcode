/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */
package main

import "fmt"

func main() {
	n := []int{1, 1, 2}
	fmt.Println(removeDuplicates(n))
	fmt.Println(n)
}

// @lc code=start
func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i += 1
		}
	}

	return len(nums)
}

// @lc code=end

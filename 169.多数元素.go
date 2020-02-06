package main

import "fmt"

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

func main() {
	fmt.Println(majorityElement([]int{1, 232, 3, 1, 342, 1, 1, 3}))
}

// @lc code=start
func majorityElement(nums []int) int {
	count := 0
	key := nums[0]

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num == key {
			count += 1
		} else {
			count -= 1
		}

		if count <= 0 {
			key = nums[i+1]
		}
	}

	return key
}

// 1 232 3 1 423 1 3

// @lc code=end

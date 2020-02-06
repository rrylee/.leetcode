/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// @lc code=start
func twoSum(nums []int, target int) []int {
	res := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := res[nums[i]]; ok {
			return []int{i, j}
		} else {
			res[target-nums[i]] = i
		}
	}
	return []int{}
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// @lc code=end

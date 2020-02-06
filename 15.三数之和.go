/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, length-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			} else if sum > 0 {
				r -= 1
			} else {
				l += 1
			}
		}
	}

	return ret
}

func threeSum1(nums []int) [][]int {
	ret := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ret
}

// @lc code=end

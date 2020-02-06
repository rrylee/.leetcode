/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedNumber := 0
	for x > reversedNumber {
		reversedNumber = reversedNumber*10 + x%10
		x /= 10
	}

	return x == reversedNumber || x == reversedNumber/10
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	nums := []int{}
	for x > 0 {
		nums = append(nums, x%10)
		x /= 10
	}

	start, end := 0, len(nums)-1
	for start < end {
		if nums[start] != nums[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	more := 1
	n := len(digits) - 1
	for i := n; i >= 0; i-- {
		sum := digits[i] + more
		if sum >= 10 {
			digits[i] = sum % 10
			more = 1
		} else {
			digits[i] = sum
			more = 0
		}
	}
	if more == 1 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}

// @lc code=end


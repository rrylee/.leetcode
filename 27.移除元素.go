/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			i += 1
		}
	}

	return len(nums)
}

// @lc code=end


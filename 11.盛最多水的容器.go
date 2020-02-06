/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	ret := 0

	start, end := 0, len(height)-1
	for start < end {
		area := (end - start) * minInt(height[start], height[end])
		if area > ret {
			ret = area
		}
		if height[start] < height[end] {
			start += 1
		} else {
			end -= 1
		}
	}

	return ret
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end


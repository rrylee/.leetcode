import "sort"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	sort.Ints(nums)
	ret := nums[0] + nums[1] + nums[2]

	for i := 0; i < length; i++ {
		l, r := i+1, length-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == target {
				return sum
			} else if sum > target {
				r -= 1
			} else {
				l += 1
			}

			if abs(sum-target) < abs(ret-target) {
				ret = sum
			}
		}
	}

	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end


import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	if length < 4 {
		return nil
	}

	ret := [][]int{}
	for a := 0; a <= length-4; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b <= length-3; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			c, d := b+1, length-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum == target {
					ret = append(ret, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c] == nums[c+1] {
						c += 1
					}
					for c < d && nums[d] == nums[d-1] {
						d -= 1
					}
					d -= 1
					c += 1
				} else if sum > target {
					d -= 1
				} else {
					c += 1
				}
			}
		}
	}
	return ret
}

// @lc code=end

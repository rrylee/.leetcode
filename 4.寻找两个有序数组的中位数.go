/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	p1, p2 := 0, 0

	isOdd := true
	if ((n1 + n2) & 1) == 0 {
		isOdd = false
	}
	middleIndex := (n1 + n2) / 2

	var cur, prev int
	for p1 < n1 || p2 < n2 {
		prev = cur

		if p1 >= n1 {
			cur = nums2[p2]
			p2 += 1
		} else if p2 >= n2 {
			cur = nums1[p1]
			p1 += 1
		} else {
			if nums1[p1] < nums2[p2] {
				cur = nums1[p1]
				p1 += 1
			} else {
				cur = nums2[p2]
				p2 += 1
			}
		}

		if p1+p2-1 == middleIndex {
			if isOdd {
				return float64(cur)
			} else {
				return (float64(cur) + float64(prev)) / 2
			}
		}
	}

	return 0
}

// @lc code=end

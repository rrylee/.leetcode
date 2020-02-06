/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

package main

import "fmt"

func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	merge(nums, 0, []int{-50, -50, -48, -47, -44, -44, -37, -35, -35, -32, -32, -31, -29, -29, -28, -26, -24, -23, -23, -21, -20, -19, -17, -15, -14, -12, -12, -11, -10, -9, -8, -5, -2, -2, 1, 1, 3, 4, 4, 7, 7, 7, 9, 10, 11, 12, 14, 16, 17, 18, 21, 21, 24, 31, 33, 34, 35, 36, 41, 41, 46, 48, 48}, 63)
	fmt.Println(nums)
}

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := len(nums1) - 1
	p2 := m - 1
	p3 := n - 1
	for p1 >= 0 {
		if p2 < 0 {
			nums1[p1] = nums2[p3]
			p3 -= 1
			p1 -= 1
			continue
		}
		if p3 < 0 {
			nums1[p1] = nums1[p2]
			p2 -= 1
			p1 -= 1
			continue
		}
		if nums1[p2] > nums2[p3] {
			nums1[p1] = nums1[p2]
			p2 -= 1
		} else {
			nums1[p1] = nums2[p3]
			p3 -= 1
		}
		p1 -= 1
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := make([]int, m, m)
	for i := 0; i < m; i++ {
		nums1Copy[i] = nums1[i]
	}

	p1 := 0
	p2 := 0
	i := 0
	for p2 < n || p1 < m {
		if p1 == m {
			nums1[i] = nums2[p2]
			p2 += 1
			i += 1
			continue
		}
		if p2 == n {
			nums1[i] = nums1Copy[p1]
			p1 += 1
			i += 1
			continue
		}
		if nums1Copy[p1] < nums2[p2] {
			nums1[i] = nums1Copy[p1]
			p1 += 1
			i += 1
			continue
		} else {
			nums1[i] = nums2[p2]
			p2 += 1
			i += 1
			continue
		}
	}
}

// @lc code=end

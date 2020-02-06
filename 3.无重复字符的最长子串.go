/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0
	left := 0

	for i := 0; i < len(s); i++ {
		if m[s[i]] > 0 && left < m[s[i]] {
			left = m[s[i]]
		}
		m[s[i]] = i + 1
		if i-left+1 > res {
			res = i - left + 1
		}
	}

	return res
}

// @lc code=end


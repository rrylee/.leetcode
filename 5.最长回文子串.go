/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

/*
if s[i] == s[j] && dp[i+1][j-1]
dp[i][j] = true
else
dp[i][j] = false
*/

package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

// @lc code=start
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := ex(s, i, i)
		len2 := ex(s, i, i+1)
		l := maxInt(len1, len2)

		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}

	return s[start : end+1]
}

func ex(s string, l, r int) (left, right, length int) {
	left, right = l, r
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return right - left + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindrome1(s string) string {
	if s == "" {
		return ""
	}

	dp := make([][]bool, len(s))
	//初始化1字母回文
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	//初始化2字母回文
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			dp[i-1][i] = true
		}
	}
	for i := 1; i < len(s); i++ {
		for j := i; j < len(s)-1; j++ {
			if s[j+1] == s[j-i] && dp[j-i+1][j] {
				dp[j-i][j+1] = true
			} else {
				dp[j-i][j+1] = false
			}
		}
	}

	max := 1
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if dp[i][j] {
				if j-i+1 > max {
					max = j - i + 1
					start = i
					end = j
				}
			}
		}
	}

	return s[start : end+1]
}

// @lc code=end

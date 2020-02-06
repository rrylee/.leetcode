/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

/*
  a b c d e
a 1 1 1 1 1
c 1 1 2 2 2
e 1 1 2 2 3
*/

/*
str1[i] != str2[j]
dp[i][j] = max(dp[i-1][j], dp[i][j-1])

str1[i] == str2[j]
dp[i][j] = dp[i-1][j-1] + 1
*/
// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxf(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func maxf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end


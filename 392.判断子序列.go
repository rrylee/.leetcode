/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

/*
t[j] == s[i]
dp[i][j] = dp[i-1][j-1]

t[j] != s[i]
dp[i][j] = dp[i][j-1]
*/
// @lc code=start
//贪心--指针
func isSubsequence(s string, t string) bool {
	sPos, tPos := 0, 0
	sLen, tLen := len(s), len(t)

	if sLen == 0 {
		return true
	}

	for tPos < tLen {
		if t[tPos] == s[sPos] {
			sPos += 1
		}

		if sPos == sLen {
			return true
		}

		tPos += 1
	}

	return false
}

//贪心--遍历数组
func isSubsequence2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for len(t) > 0 {
		a := t[len(t)-1]
		t = t[0 : len(t)-1]
		if a == s[len(s)-1] {
			s = s[0 : len(s)-1]
			if len(s) == 0 {
				return true
			}
		}
	}
	return false
}

//动态规划
func isSubsequence1(s string, t string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	m := len(t)
	if m == 0 {
		return false
	}

	dp := [][]bool{}
	for i := 0; i <= n; i++ {
		line := make([]bool, m+1)
		if i == 0 {
			for ii, _ := range line {
				line[ii] = true
			}
		}
		dp = append(dp, line)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if t[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[n][m]
}

// @lc code=end

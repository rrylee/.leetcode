/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

/*
s[i] == p[j] || p[j] == '?'
dp[i][j] = dp[i-1]][j-1]

p[i] == '*'
dp[i][j] = dp[i-1][j] || dp[i][j-1]
*/
/*
    a b c d e d e
  1 0 0 0 0 0 0 0
a 0 1 1 1 1 1 1 1
* 0 1 1 1 1 1 1 1
? 0 0 1 1 1 1 1 1
d 0 0 0 0 1 0 1 0
e 0 0 0 0 0 1 0 1
e 0 0 0 0 0 1 0 0
*/

// @lc code=start
func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}

	m, n := len(s), len(p)

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] != '*' {
			break
		} else {
			dp[i][0] = true
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[j-1] == p[i-1] || p[i-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}

// @lc code=end

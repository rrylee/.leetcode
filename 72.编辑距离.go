import "fmt"

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

/*
  h o r s e
r 1 2 3 4 5
o 2 1 2 3 4
s 3 2 3 2 3
*/

/*
word1[i]==word2[j]
dp[i][j] = dp[i-1][j-1]

word1[u]!=work2[j]
dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
*/

// @lc code=start
func minDistance(word1 string, word2 string) int {
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}
	m, n := len(word1), len(word2)
	dp := make([][]int, n+1) //[n][m]

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}

	for _, line := range dp {
		fmt.Println(line)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[j-1] != word2[i-1] {
				dp[i][j] = minf(dp[i-1][j], minf(dp[i-1][j-1], dp[i][j-1])) + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}

func minf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

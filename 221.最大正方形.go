/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

/*
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

*/

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	dp := [][]int{}

	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					tmp := dp[i-1][j]
					if tmp > dp[i-1][j-1] {
						tmp = dp[i-1][j-1]
					}
					if tmp > dp[i][j-1] {
						tmp = dp[i][j-1]
					}
					dp[i][j] = tmp + 1
				}
			}

			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max * max
}

// @lc code=end


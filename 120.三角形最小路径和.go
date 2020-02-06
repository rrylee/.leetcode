/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	height := len(triangle)
	if height == 1 {
		return triangle[0][0]
	}

	for i := height - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			tmp := triangle[i+1][j]
			if tmp > triangle[i+1][j+1] {
				tmp = triangle[i+1][j+1]
			}
			triangle[i][j] += tmp
		}
	}

	return triangle[0][0]
}

// @lc code=end

func minimumTotal1(triangle [][]int) int {
	height := len(triangle)
	width := len(triangle[height-1])
	dp := [][]int{}

	for i := 0; i < height; i++ {
		dp = append(dp, make([]int, width))
	}

	lastRow := triangle[height-1]
	for i := 0; i < width; i++ {
		dp[height-1][i] = lastRow[i]
	}

	for i := height - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			one := dp[i+1][j]
			if one > dp[i+1][j+1] {
				one = dp[i+1][j+1]
			}
			dp[i][j] = one + triangle[i][j]
		}
	}
	return dp[0][0]
}

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				if obstacleGrid[0][j] == 1 {
					for k := j; k < width; k++ {
						obstacleGrid[0][k] = 0
					}
					break
				} else {
					obstacleGrid[0][j] = 1
					continue
				}
			}
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				if j == 0 {
					if obstacleGrid[i][j] == 1 {
						obstacleGrid[i][j] = 0
					} else {
						obstacleGrid[i][j] = obstacleGrid[i-1][j]
					}
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
				}
			}
		}
	}
	return obstacleGrid[height-1][width-1]
}

// @lc code=end


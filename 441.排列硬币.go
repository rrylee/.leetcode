/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	// (1 + x)*x/2 = n
	sum := 0
	i := 0
	for sum < n {
		i += 1
		sum += i
	}
	if sum == n {
		return i
	} else {
		return i - 1
	}
}

// @lc code=end

x(1+x)/2=n
x/2+x*x/2=n
(x/2+1/2)x=n

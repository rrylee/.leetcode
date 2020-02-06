/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

/*
s[i] == ')' && s[i-1] == '('
dp[i] = dp[i-2] + 2

s[i] == ')' && s[i-1] == ')' && s[i-1-dp[i-1]] == '('
dp[i] = dp[i-1] + dp[i - dp[i-1] - 2] + 2

*/
package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("(()())"))
}

// @lc code=start
func longestValidParentheses(s string) int {
	dp := make([]int, len(s), len(s))
	max := 0

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else if i-dp[i-1] > 0 && s[i-1-dp[i-1]] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-2-dp[i-1]] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}

			if dp[i] > max {
				max = dp[i]
			}
		}
	}

	return max
}

// @lc code=end

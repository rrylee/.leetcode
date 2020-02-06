/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"c", "c"}))
}

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	if strs[0] == "" {
		return ""
	}

	i := 0
	isBreak := false
	for {
		if isBreak || i == len(strs[0]) {
			break
		}
		cur := strs[0][i]
		for j, str := range strs {
			if j == 0 {
				continue
			}
			if len(str) <= i {
				isBreak = true
				break
			}
			if cur != str[i] {
				isBreak = true
				break
			}
		}
		if !isBreak {
			i += 1
		}
	}

	return strs[0][0:i]
}

// @lc code=end

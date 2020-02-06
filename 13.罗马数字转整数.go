/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
package main

import "fmt"

func main() {
	fmt.Println((romanToInt("III")))
}

// @lc code=start
func romanToInt(s string) int {
	m := map[byte]int{
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ret := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			if i+1 < n {
				if s[i+1] == 'V' {
					ret += 4
					i += 1
				} else if s[i+1] == 'X' {
					ret += 9
					i += 1
				} else {
					ret += 1
				}
			} else {
				ret += 1
			}
			continue
		}
		if s[i] == 'X' {
			if i+1 < n {
				if s[i+1] == 'L' {
					ret += 40
					i += 1
				} else if s[i+1] == 'C' {
					ret += 90
					i += 1
				} else {
					ret += 10
				}
			} else {
				ret += 10
			}
			continue
		}
		if s[i] == 'C' {
			if i+1 < n {
				if s[i+1] == 'D' {
					ret += 400
					i += 1
				} else if s[i+1] == 'M' {
					ret += 900
					i += 1
				} else {
					ret += 100
				}
			} else {
				ret += 100
			}
			continue
		}
		ret += m[s[i]]
	}
	return ret
}

// @lc code=end

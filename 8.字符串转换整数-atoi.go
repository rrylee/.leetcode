/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

package main

import "fmt"

func main() {
	fmt.Println(myAtoi("+-2"))
}

// @lc code=start
func myAtoi(str string) int {
	isNev := false
	checkFirst := true
	res := 0

	bToi := map[uint8]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	for i := 0; i < len(str); i++ {
		if checkFirst && (str[i] == ' ') {
			fmt.Println("-000")
			continue
		}
		if checkFirst && str[i] == '-' {
			checkFirst = false
			isNev = true
			continue
		}
		if checkFirst && str[i] == '+' {
			checkFirst = false
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			break
		}

		if str[i] >= '0' || str[i] <= '9' {
			fmt.Println(res, bToi[str[i]])
			res = res*10 + bToi[str[i]]
			if checkFirst {
				checkFirst = false
			}
			if !isNev && res > 2147483647 {
				return 2147483647
			}
			if isNev && (-res) < -2147483648 {
				return -2147483648
			}
			continue
		}
	}

	if isNev {
		return -res
	}

	return res
}

// @lc code=end

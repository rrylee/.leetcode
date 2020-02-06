/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
package main

import "fmt"

func main() {
	fmt.Println(intToRoman(20))
}

// @lc code=start
func intToRoman(num int) string {
	mm := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	m := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	roman := ""
	for num > 0 {
		for _, i := range mm {
			if num/i > 0 {
				roman += m[i]
				num -= i
				break
			}
		}
	}
	return roman
}

// @lc code=end

package main

import "fmt"

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

func main() {
	fmt.Println(myPow(2.00000, 10))
}

// 1  2  4
// 2, 4, 16

// @lc code=start
func myPow(x float64, n int) float64 {
	m := map[int]float64{
		1: x,
	}

	isNev := false
	if n < 0 {
		isNev = true
		n = -n
	}

	cur := 1
	var ret float64 = 1
	for n > 0 {
		if n >= cur {
			v := m[cur]
			ret = ret * v
			n -= cur
			cur *= 2
			m[cur] = v * v
		} else {
			cur /= 2
		}
	}

	if !isNev {
		return ret
	}
	return 1 / ret
}

// @lc code=end

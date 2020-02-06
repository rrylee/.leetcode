package main

import "fmt"

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

// @lc code=start
func addBinary(a string, b string) string {
	ret := ""

	p1 := len(a) - 1
	p2 := len(b) - 1

	up := 0
	for p1 >= 0 || p2 >= 0 {
		aa, bb := 0, 0
		if p1 < 0 {
			bb = getI(b[p2])
		} else if p2 < 0 {
			aa = getI(a[p1])
		} else {
			aa = getI(a[p1])
			bb = getI(b[p2])
		}

		switch aa + bb + up {
		case 3:
			ret = "1" + ret
			up = 1
			break
		case 2:
			ret = "0" + ret
			up = 1
			break
		case 1:
			ret = "1" + ret
			up = 0
			break
		case 0:
			ret = "0" + ret
			up = 0
			break
		}
		p1 -= 1
		p2 -= 1
	}
	if up == 1 {
		ret = "1" + ret
	}

	return ret
}

func getI(k byte) int {
	return int(uint8(k - '0'))
}

// @lc code=end

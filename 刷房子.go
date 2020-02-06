package main

import "fmt"

func main() {
	fmt.Println(minMoney([][]int{
		{7, 2, 9},
		{5, 13, 8},
		{9, 2, 4},
	}))
}

func minMoney(nums [][]int) int {
	l := len(nums)
	min := 0
	for i := 0; i < l; i++ {
		if i == 0 {
			min = minf(nums[0][0], nums[0][1], nums[0][2])
			continue
		}

		nums[i][0] += minf(nums[i-1][1], nums[i-1][2])
		nums[i][1] += minf(nums[i-1][0], nums[i-1][2])
		nums[i][2] += minf(nums[i-1][0], nums[i-1][1])

		if i == l-1 {
			min = minf(nums[i][0], nums[i][1], nums[i][2])
		}
	}
	fmt.Println(nums)
	return min
}

func minf(nums ...int) int {
	n := nums[0]

	for _, num := range nums {
		if num < n {
			n = num
		}
	}

	return n
}

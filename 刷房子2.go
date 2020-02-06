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
	m := len(nums)
	n := len(nums[0])

	for i := 1; i < m; i++ {
		min1, min2, minIndex := getMin1Min2AndMin1Index(nums[i-1])
		for j := 0; j < n; j++ {
			if j == minIndex {
				nums[i][j] += min2
			} else {
				nums[i][j] += min1
			}
		}
	}

	min := nums[m-1][0]
	for i := 0; i < n; i++ {
		min = minf(min, nums[m-1][i])
	}

	return min
}

func getMin1Min2AndMin1Index(nums []int) (int, int, int) {
	min1 := nums[0]
	min2 := nums[1]
	minIndex := 0
	if min1 > min2 {
		min1, min2 = min2, min1
		minIndex = 1
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] < min1 {
			min1, min2, minIndex = min2, nums[i], i
		} else if nums[i] > min2 {
			continue
		} else {
			min2 = nums[i]
		}
	}
	return min1, min2, minIndex
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

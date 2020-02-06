// type Times struct {
// 	count int
// 	index int
// }

// type SortBy []Times

// func (a SortBy) Len() int           { return len(a) }
// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}

func minSetSize(arr []int) int {
	m := map[int]int{}

	for _, i := range arr {
		m[i] += 1
	}

	times := []int{}
	for _, v := range m {
		times = append(times, v)
	}
	sort.Ints(times)

	target := len(arr) / 2
	if len(arr)&1 == 1 {
		target += 1
	}

	count := 0
	j := 0
	for i := len(times) - 1; i >= 0; i-- {
		j += 1
		count += times[i]
		if count >= target {
			break
		}
	}
	return j
}

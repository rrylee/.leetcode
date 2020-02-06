package main

import (
	"fmt"
	"sort"
)

func main() {
	m := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(kWeakestRows(m, 3))
}

func kWeakestRows(mat [][]int, k int) []int {
	m := map[int][]int{}

	for i, nums := range mat {
		soliders := 0
		for _, num := range nums {
			if num == 0 {
				break
			}
			soliders += 1
		}
		if _, ok := m[soliders]; ok {
			m[soliders] = append(m[soliders], i)
		} else {
			m[soliders] = []int{i}
		}
	}

	sortKeys := []int{}
	for k, _ := range m {
		sortKeys = append(sortKeys, k)
	}
	sort.Ints(sortKeys)

	ret := []int{}
	i := 0
	for {
		indexes := m[sortKeys[i]]
		if len(indexes) == 1 {
			ret = append(ret, indexes[0])
		} else {
			sort.Ints(indexes)
			for j := 0; j < len(indexes); j++ {
				ret = append(ret, indexes[j])
				if len(ret) == k {
					break
				}
			}
		}

		if len(ret) == k {
			break
		}

		i += 1
	}

	return ret
}

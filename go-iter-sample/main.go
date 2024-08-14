package main

import (
	"fmt"
	"maps"
	"slices"
)

type taple struct {
	a int
	b string
}

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range slices.Chunk(num, 3) {
		fmt.Println(i)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range maps.All(m) {
		fmt.Println(k, v)
	}

	// random slice order
	arr := []int{3, 1, 2, 4, 5}
	fmt.Println(slices.IsSorted(arr))
	slices.SortFunc(arr, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	fmt.Println(arr)

	// sort の種類を設定してみる
	taples1 := []taple{{3, "dd"}, {1, "a"}, {3, "cc"}, {3, "ff"}, {2, "b"}, {4, "ee"}}
	s1 := slices.SortedFunc(slices.Values(taples1), func(a, b taple) int {
		return a.a - b.a
	})
	fmt.Println(s1)

	taples2 := []taple{{3, "dd"}, {1, "a"}, {3, "cc"}, {3, "ff"}, {2, "b"}, {4, "ee"}}
	s2 := slices.SortedStableFunc(slices.Values(taples2), func(a, b taple) int {
		return a.a - b.a
	})
	fmt.Println(s2)
}

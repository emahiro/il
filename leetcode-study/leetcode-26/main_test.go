package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func uniqNormal(s *[]int) {
	m := make(map[int]struct{})
	for _, v := range *s {
		m[v] = struct{}{}
	}
	*s = (*s)[:0]
	for k := range m {
		*s = append(*s, k)
	}
}

func uniqFast(s *[]int) {
	offset := 1
	for i := 1; i < len(*s); i++ {
		if (*s)[i] != (*s)[i-1] {
			(*s)[offset] = (*s)[i]
			offset++
		}
	}
	*s = (*s)[:offset]
}

func uniqGen[T comparable](s *[]T) {
	_ = slices.Compact(*s)
}

func Benchmark(b *testing.B) {
	b.ReportAllocs()

	list := []int{1, 1, 2, 2}

	b.Run("unique normal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			uniqNormal(&list)
		}
	})

	b.Run("unique faster", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			uniqFast(&list)
		}
	})

	b.Run("unique generic", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			uniqGen(&list)
		}
	})
}

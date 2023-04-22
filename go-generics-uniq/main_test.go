package main

import "testing"

func BenchmarkUniq(b *testing.B) {
	s := []string{"1", "2", "3", "10", "9", "10", "9", "8", "100", "200", "100"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Uniq(s)
	}
}

func BenchmarkPureUniq(b *testing.B) {
	s := []string{"1", "2", "3", "10", "9", "10", "9", "8", "100", "200", "100"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PureUniq(s)
	}
}

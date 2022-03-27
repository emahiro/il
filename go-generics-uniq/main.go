package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	intMap := map[string]int64{
		"1": 1,
		"2": 2,
	}
	a := SumInts(intMap)
	fmt.Println(a)

	floatMap := map[string]float64{
		"1": 1.1,
		"2": 2.2,
	}
	b := SumFloat(floatMap)
	fmt.Println(b)

	stringMap := map[string]string{
		"1": "a",
		"2": "b",
	}
	c := SumString(stringMap)
	fmt.Println(c)

	fmt.Println("Generic implementation is below")

	aa := SumT(intMap)
	fmt.Println(aa)
	bb := SumT(floatMap)
	fmt.Println(bb)
	cc := SumT(stringMap)
	fmt.Println(cc)

	fmt.Println("more generic below")
	m1 := map[int]int64{
		1: 1,
		2: 2,
	}
	aaa := Sum(m1)
	fmt.Println(aaa)

	m2 := map[float32]float64{
		1.1: 1.1,
		2.2: 2.2,
	}
	bbb := Sum(m2)
	fmt.Println(bbb)

	fmt.Println("get keys process below")
	kv1 := map[string]interface{}{
		"1": struct{}{},
		"2": struct{}{},
		"3": struct{}{},
	}
	fmt.Println(PureKeys(kv1))

	kv2 := map[int]interface{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
	}
	fmt.Println(Keys(kv1))
	fmt.Println(Keys(kv2))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloat(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumString(m map[string]string) string {
	var s string
	for _, v := range m {
		s += v
	}
	return s
}

func SumT[T int64 | float64 | string](m map[string]T) T {
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

func Sum[K comparable, T int64 | float64 | string](m map[K]T) T {
	var t T
	for _, v := range m {
		t += v
	}
	return t
}

func PureKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Keys[T comparable](m map[T]interface{}) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Uniq[T comparable](s []T) []T {
	slices.Sort(s)
	return slices.Clip(slices.Compact(s))
}

func PureUniq(s []string) []string {
	var ss []string
	m := make(map[string]struct{}, len(s))
	for _, ss := range s {
		m[ss] = struct{}{}
	}
	for k := range m {
		ss = append(ss, k)
	}
	return ss
}

package main

import "fmt"

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

package main

import "testing"

// func BenchmarkFib10(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		Fib(10)
// 	}
// }
var result int

func benchmarkFib(i int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Fib(i)
	}
	result = r
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib30(b *testing.B) { benchmarkFib(30, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

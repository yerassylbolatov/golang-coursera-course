package main

import (
	"testing"
)

const iter = 10

func BenchmarkEmptyAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []int{}
		for j := 0; j < 10; j++ {
			data = append(data, j)
		}
	}
}

func BenchmarkPreAllocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, 0, iter)
		for j := range data {
			data[j] = j
		}
	}
}

func BenchmarkPreAllocateWithZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, iter)
		for j := range data {
			data[j] = j
		}
	}
}

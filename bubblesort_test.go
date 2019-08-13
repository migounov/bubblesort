package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Int())
	}
	return s
}

func BenchmarkBubblesort_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := generateSlice(10)
		b.StartTimer()
		bubbleSort(s)
	}
}

func BenchmarkBubblesort_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := generateSlice(100)
		b.StartTimer()
		bubbleSort(s)
	}
}

func BenchmarkBubblesort_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := generateSlice(1000)
		b.StartTimer()
		bubbleSort(s)
	}
}

func BenchmarkBubblesort_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := generateSlice(10000)
		b.StartTimer()
		bubbleSort(s)
	}
}

func BenchmarkBubblesort_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := generateSlice(100000)
		b.StartTimer()
		bubbleSort(s)
	}
}

func TestBubblesort_10(t *testing.T) {
	s := generateSlice(10)
	e := make([]int, 10)
	copy(e, s)
	sort.Ints(e)
	bubbleSort(s)
	if !reflect.DeepEqual(s, e) {
		fmt.Printf("Our sorted array: %v\n", s)
		fmt.Printf("Expected array:   %v\n", e)
		t.Errorf("Array is not sorted!")
	}
}


package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Int())
	}
	return s
}

func BenchmarkBubblesort(b *testing.B) {
	type parameters struct{
		name		string
        sliceSize 	int
	}

	var params []parameters
	for i := 1; i <= 10; i++ {
		params = append(params, parameters{strconv.Itoa(i*10000), i*10000})
	}

	for _, p := range params {
		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				s := generateSlice(p.sliceSize)
				b.StartTimer()
				bubbleSort(s)
			}
		})
	}
}

func TestBubblesort(t *testing.T) {
	s := generateSlice(100)
	e := make([]int, 100)
	copy(e, s)
	sort.Ints(e)
	bubbleSort(s)
	if !reflect.DeepEqual(s, e) {
		fmt.Printf("Our sorted array: %v\n", s)
		fmt.Printf("Expected array:   %v\n", e)
		t.Errorf("Array is not sorted!")
	}
}


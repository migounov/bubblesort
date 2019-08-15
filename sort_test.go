package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

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
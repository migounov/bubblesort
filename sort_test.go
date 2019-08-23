package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type sortMethod func(s []int) []int

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Int())
	}
	return s
}

func TestSort(t *testing.T) {
	type parameters struct {
		name      	string
		sliceSize 	int
		algorithm	sortMethod
	}

	var params []parameters
	params = append(params, parameters{"Bubble", 1, bubbleSort})
	params = append(params, parameters{"Bubble", 1000, bubbleSort})
	params = append(params, parameters{"Merge", 1, mergeSort})
	params = append(params, parameters{"Merge", 1000, mergeSort})
	params = append(params, parameters{"Quick", 1, quickSort})
	params = append(params, parameters{"Quick", 1000, quickSort})

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			s := generateSlice(p.sliceSize)
			act := make([]int, 0)
			exp := make([]int, p.sliceSize)
			copy(exp, s)
			sort.Ints(exp)

			act = p.algorithm(s)

			if !reflect.DeepEqual(act, exp) {
				fmt.Printf("Our sorted array: %v\n", act)
				fmt.Printf("Expected array:   %v\n", exp)
				t.Errorf("Array is not sorted!")
			}
		})
	}
}

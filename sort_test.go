package sort

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

func TestSort(t *testing.T) {
	type parameters struct {
		name      string
		sliceSize int
	}

	var params []parameters
	params = append(params, parameters{"Bubble", 1})
	params = append(params, parameters{"Bubble", 1000})
	params = append(params, parameters{"Merge", 1})
	params = append(params, parameters{"Merge", 1000})
	params = append(params, parameters{"Quick", 1})
	params = append(params, parameters{"Quick", 1000})

	for _, p := range params {
		t.Run(p.name, func(t *testing.T) {
			s := generateSlice(p.sliceSize)
			act := make([]int, 0)
			exp := make([]int, p.sliceSize)
			copy(exp, s)
			sort.Ints(exp)

			switch p.name {
			case "Bubble":
				bubbleSort(s)
				act = s
			case "Merge":
				act = mergeSort(s)
			case "Quick":
				act = quickSort(s)
			}

			if !reflect.DeepEqual(act, exp) {
				fmt.Printf("Our sorted array: %v\n", act)
				fmt.Printf("Expected array:   %v\n", exp)
				t.Errorf("Array is not sorted!")
			}
		})
	}
}

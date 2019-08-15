package main

import (
	"fmt"
	"math/rand"
)

func generateSlice(n int) []int {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Int())
	}
	return s
}

func bubbleSort(s []int) {
	for countPasses, swapped := 1, true; countPasses < len(s) && swapped; countPasses++ {
		swapped = false
		for j := 0; j < len(s)-countPasses; j++ {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
				swapped = true
			}
		}
	}
}

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

func main() {
	s := generateSlice(10)
	fmt.Printf("%v\n%v\n", s, MergeSort(s))
}
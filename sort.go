package sort

import (
	"math/rand"
)

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

func merge(left, right []int) []int {
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

func mergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := mergeSort(s[:n])
	r := mergeSort(s[n:])
	return merge(l, r)
}

func quickSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	insertSpot := 0
	pivot := rand.Int() % len(s)
	s[pivot], s[len(s)-1] = s[len(s)-1], s[pivot]
	pivot = len(s) - 1

	for i := 0; i < pivot; i++ {
		if s[i] < s[pivot] {
			s[insertSpot], s[i] = s[i], s[insertSpot]
			insertSpot++
		}
	}

	s[insertSpot], s[pivot] = s[pivot], s[insertSpot]
	quickSort(s[:insertSpot])
	quickSort(s[insertSpot+1:])

	return s
}

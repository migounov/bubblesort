package main

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

func main() {
}

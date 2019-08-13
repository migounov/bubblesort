package main

func bubbleSort(s []int) {
	for count_passes, swapped := 1, true; count_passes < len(s) && swapped; count_passes++ {
		swapped = false
		for j := 0; j < len(s) - count_passes; j++ {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
				swapped = true
			}
		}
	}
}

func main() {
}
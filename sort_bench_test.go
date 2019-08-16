package sort

import (
	"strconv"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	type parameters struct {
		algorithm string
		name      string
		sliceSize int
	}

	var params []parameters
	/*for i := 1; i <= 50; i++ {
		params = append(params, parameters{"Bubble", strconv.Itoa(i * 2) + "K", i * 2000})
	}
	for i := 1; i <= 50; i++ {
		params = append(params, parameters{"Merge", strconv.Itoa(i * 2) + "M", i * 2000000})
	}*/
	for i := 1; i <= 50; i++ {
		params = append(params, parameters{"Quick", strconv.Itoa(i*2) + "M", i * 2000000})
	}

	for _, p := range params {
		b.Run(p.algorithm+p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := generateSlice(p.sliceSize)
				b.StartTimer()
				switch p.algorithm {
				case "Bubble":
					bubbleSort(s)
				case "Merge":
					mergeSort(s)
				case "Quick":
					quickSort(s)
				}
				b.StopTimer()
			}
		})
	}
}

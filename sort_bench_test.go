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
	for _, alg := range []string{"Bubble", "Merge", "Quick"} {
		for i := 1; i <= 10; i++ {
			params = append(params, parameters{alg, strconv.Itoa(i) + "K", i * 1000})
		}
		for i := 2; i <= 10; i++ {
			params = append(params, parameters{alg, strconv.Itoa(i*10) + "K", i * 10000})
		}
		if alg != "Bubble" {
			for i := 2; i <= 10; i++ {
				params = append(params, parameters{alg, strconv.Itoa(i*100) + "K", i * 100000})
			}
			for i := 2; i <= 10; i++ {
				params = append(params, parameters{alg, strconv.Itoa(i) + "M", i * 1000000})
			}
			for i := 2; i <= 10; i++ {
				params = append(params, parameters{alg, strconv.Itoa(i*10) + "M", i * 10000000})
			}
		}
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

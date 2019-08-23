package sort

import (
	"strconv"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	type parameters struct {
		name      string
		sliceSize int
		algorithm sortMethod
	}

	var params []parameters
	for i := 1; i <= 10; i++ {
		params = append(params, parameters{"Bubble" + strconv.Itoa(i) + "K", i * 1000, bubbleSort})
	}
	for i := 2; i <= 10; i++ {
		params = append(params, parameters{"Bubble" + strconv.Itoa(i*10) + "K", i * 10000, bubbleSort})
	}
	for i := 2; i <= 10; i++ {
		params = append(params, parameters{"Merge" + strconv.Itoa(i*100) + "K", i * 100000, mergeSort})
	}
	for i := 2; i <= 10; i++ {
		params = append(params, parameters{"Merge" + strconv.Itoa(i) + "M", i * 1000000, mergeSort})
	}
	for i := 2; i <= 10; i++ {
		params = append(params, parameters{"Merge" + strconv.Itoa(i*10) + "M", i * 10000000, mergeSort})
	}
	for i := 2; i <= 10; i++ {
		params = append(params, parameters{"Quick" + strconv.Itoa(i*100) + "K", i * 100000, quickSort})
	}
	for i := 2; i <= 10; i++ {
		params = append(params, parameters{"Quick" + strconv.Itoa(i) + "M", i * 1000000, quickSort})
	}
	for i := 2; i <= 10; i++ {
		params = append(params, parameters{"Quick" + strconv.Itoa(i*10) + "M", i * 10000000, quickSort})
	}

	for _, p := range params {
		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := generateSlice(p.sliceSize)
				b.StartTimer()
				p.algorithm(s)
				b.StopTimer()
			}
		})
	}
}

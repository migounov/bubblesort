package main

import (
	"strconv"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	type parameters struct {
		name      string
		sliceSize int
	}

	var params []parameters
	for i := 1; i <= 50; i++ {
		params = append(params, parameters{strconv.Itoa(i * 2000), i * 2000})
	}

	for _, p := range params {
		b.Run(p.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				s := generateSlice(p.sliceSize)
				b.StartTimer()
				bubbleSort(s)
			}
		})
	}
}
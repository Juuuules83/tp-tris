package main

import (
	"fmt"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n) // préparation hors de la mesure
		scores := make([]int, n)
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base) // chaque tour repart du même désordre
				b.StartTimer()
				BubbleSort(scores)
			}
		})
	}
}

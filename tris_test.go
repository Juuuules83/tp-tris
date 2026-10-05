package main

import (
	"fmt"
	"testing"
)

// Benchmark demandé dans la mise en place du TP.
// Le copy permet de repartir du même désordre à chaque tour.
func BenchmarkBubbleSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				BubbleSort(scores)
			}
		})
	}
}


func BenchmarkBubbleSortSorted(b *testing.B) {
	base := SortedScores(10_000)
	scores := make([]int, len(base))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(scores, base)
		b.StartTimer()

		BubbleSort(scores)
	}
}


func BenchmarkBubbleSortReversed(b *testing.B) {
	base := ReversedScores(10_000)
	scores := make([]int, len(base))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(scores, base)
		b.StartTimer()

		BubbleSort(scores)
	}
}

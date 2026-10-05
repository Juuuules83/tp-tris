package main

func BubbleSort(scores []int) {
	for {
		echange := false

		for i := 0; i < len(scores)-1; i++ {
			if scores[i] > scores[i+1] {
				scores[i], scores[i+1] = scores[i+1], scores[i]
				echange = true
			}
		}

		if !echange {
			break
		}
	}
}
package main

// BubbleSort trie scores dans l'ordre croissant.
func BubbleSort(scores []int) {
	for lastIndex := len(scores) - 1; lastIndex > 0; lastIndex-- {
		swapped := false

		for i := 0; i < lastIndex; i++ {
			if scores[i] > scores[i+1] {
				scores[i], scores[i+1] = scores[i+1], scores[i]
				swapped = true
			}
		}

		// Si aucun échange n'a eu lieu, le slice est déjà trié.
		if !swapped {
			return
		}
	}
}

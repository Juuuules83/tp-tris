package main

import "fmt"

func main() {
	liste := []int{5, 2, 8, 1, 7, 3, 6, 4}

	fmt.Println("avant :", liste)
	BubbleSort(liste)
	fmt.Println("après :", liste)
}

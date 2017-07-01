package main

import (
	"fmt"
)

func main() {
	a, b := readData()
	ta, tb := sort(a, b)
	fmt.Printf("%d %d", ta, tb)
}

func sort(a []int, b []int) (int, int) {
	aTotal := 0
	bTotal := 0

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			aTotal += 1
		} else if a[i] < b[i] {
			bTotal += 1
		}
	}
	return aTotal, bTotal
}

func readData() ([]int, []int) {

	a := make([]int, 3)
	b := make([]int, 3)

	for i := range a {
		fmt.Scanf("%d", &a[i])
	}

	for i := range b {
		fmt.Scanf("%d", &b[i])
	}

	return a, b
}

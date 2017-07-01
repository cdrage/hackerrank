package main

import (
	"fmt"
)

func main() {
	a := readData()
	fmt.Printf("%d", sum(a))
}

func sum(a []int64) int64 {
	var sum int64
	for i := range a {
		sum += a[i]
	}
	return sum
}

func readData() []int64 {

	var size int64
	fmt.Scanf("%d", &size)

	a := make([]int64, size)

	for i := range a {
		fmt.Scanf("%d", &a[i])
	}

	return a
}

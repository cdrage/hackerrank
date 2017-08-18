package main

import (
	"fmt"
	"log"
)

func main() {
	a := readData()

	first := 0
	second := 0
	length := len(a)

	for i, j := 0, length; i < length; i++ {
		j--
		first += a[i][i]
		second += a[i][j]
	}

	if first > second {
		fmt.Println(first - second)
	} else {
		fmt.Println(second - first)
	}
}

func readData() [][]int {

	var size int
	fmt.Scanf("%d", &size)

	a := make([][]int, size)

	for i := range a {
		a[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			_, err := fmt.Scanf("%d", &a[i][j])
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return a
}

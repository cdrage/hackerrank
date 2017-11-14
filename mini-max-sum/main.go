package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	min, max := calculateMinMax(data)
	fmt.Printf("%d %d\n", min, max)
}

func calculateMinMax(array []int) (int, int) {

	sum := 0
	min := 0
	max := 0

	for _, i := range array {
		sum += i
		if i > max {
			max = i
		} else if i < min {
			min = i
		}
	}

	sort.Ints(array)

	return sum - array[len(array)-1], sum - array[0]
}

func readData() ([]int, error) {

	data := make([]int, 5)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

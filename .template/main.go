package main

import (
	"fmt"
	"log"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d", data)
	//min, max := calculateMinMax(data)
	//fmt.Printf("%d %d\n", min, max)
}

func calculateMinMax(array []int) int {
	return 0
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

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	data, rotation, err := readData()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%d %d", data, rotation)

	// left rotation of n shifts,
	output := rotate(data, rotation)
	fmt.Printf(strings.Trim(fmt.Sprint(output), "[]"))

}

func rotate(array []int, rotate int) []int {

	new_array := make([]int, len(array))

	for j := 0; j < len(array); j++ {

		location := (j + rotate) % len(array)
		new_array[j] = array[location]
	}

	return new_array
}

func readData() ([]int, int, error) {

	var size int
	fmt.Scanf("%d", &size)

	var rotation int
	fmt.Scanf("%d", &rotation)

	data := make([]int, size)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, 0, err
		}
	}

	return data, rotation, nil
}

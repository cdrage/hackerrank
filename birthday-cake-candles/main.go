package main

import (
	"fmt"
	"log"
)

/*
	4 years old, cake will have:
	4 candles of height 3, 2, 1, 3
	Can only blow out the tallest ones (3 and 3)
*/

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Age: %d\nCandles %d\n", len(data), data)
	//fmt.Printf("Amount of candles you can blow: %d", birthdayCakeCandles(data))
	fmt.Printf("%d", birthdayCakeCandles(data))
	//min, max := calculateMinMax(data)
	//fmt.Printf("%d %d\n", min, max)
}

func birthdayCakeCandles(array []int) int {

	candles := 0
	tallest := 0

	for _, i := range array {
		if i > tallest {
			tallest = i
			candles = 1
		} else if i == tallest {
			candles = candles + 1
		}
	}

	return candles
}

func readData() ([]int, error) {

	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int, length)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

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

	p, m, z := plusMinusZero(data)
	fmt.Printf("%f\n%f\n%f", p, m, z)
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

func plusMinusZero(nums []int) (float64, float64, float64) {

	var plus, minus, zero float64
	length := float64(len(nums))

	for _, v := range nums {
		if v > 0 {
			plus++
		} else if v < 0 {
			minus++
		} else {
			zero++
		}

	}

	return plus / length, minus / length, zero / length
}

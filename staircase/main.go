package main

import (
	"fmt"
	"strings"
)

func main() {
	var length int
	fmt.Scanf("%d", &length)
	stairs(length)
}

func stairs(height int) {
	for i := 1; i <= height; i++ {
		fmt.Printf("%s%s\n",
			strings.Repeat(" ", height-i),
			strings.Repeat("#", i))
	}
}

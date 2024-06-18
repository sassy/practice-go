package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func sqrt(x int) float64 {
	return math.Sqrt(float64(x))
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Printf("Now you have %g problems.\n", sqrt(7))
}

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

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Printf("Now you have %g problems.\n", sqrt(7))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
}

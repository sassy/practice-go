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

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

func main() {
	k := 3
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	fmt.Println(add(42, 13))
	fmt.Printf("Now you have %g problems.\n", sqrt(7))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

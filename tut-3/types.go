package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

func main() {
	var num1, num2 float64 = 5.5, 7.4
	fmt.Println(add(num1, num2))
}

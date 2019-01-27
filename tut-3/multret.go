package main

import (
	"fmt"
)

func multiple(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := "hey", "pk"
	fmt.Println(multiple(a, b))
}

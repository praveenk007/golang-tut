package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]float32)

	grades["timmy"] = 32
	grades["pk"] = 40
	fmt.Println(grades)
}

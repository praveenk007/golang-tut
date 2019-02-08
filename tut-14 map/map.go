package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]float32)

	grades["timmy"] = 32
	grades["pk"] = 40
	fmt.Println(grades)
	delete(grades, "timmy")
	fmt.Println(grades)

	iterateMap(grades)
}

func iterateMap(grades map[string]float32) {
	for k, v := range grades {
		fmt.Println("k: ", k, ", v : ", v)
	}
}

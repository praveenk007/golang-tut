package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== basic ===")
	basicFor()
	fmt.Println("=== basic ===")
	fmt.Println("=== enhanced ===")
	enhancedFor()
	fmt.Println("=== enhanced ===")
	/*fmt.Println("=== infinite ===")
	infiniteFor()
	fmt.Println("=== infinite ===")*/
	fmt.Println("=== break ===")
	breakFor()
	fmt.Println("=== break ===")
}

func basicFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func enhancedFor() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func breakFor() {
	i := 2
	for {
		i += 4
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
}

/*func infiniteFor() {
	for {
		fmt.Println("Do stuff")
	}
}*/

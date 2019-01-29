package main

import "fmt"

func main() {
	x := 10
	a := &x
	fmt.Println(a)
	fmt.Println(*a)
	*a = 2
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
}

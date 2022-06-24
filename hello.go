package main

import (
	"fmt"
)

var sayi = 5

func add(x int, y int) int {
	return x + y
}

func multiple(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(5, 10))
	fmt.Println(multiple(5, 10))
}

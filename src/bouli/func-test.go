package main

import (
	"fmt"
)

func main() {
	f := closure(1)
	fmt.Println(f(1))
	fmt.Println(f(1))
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

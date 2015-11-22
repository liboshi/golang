package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	sum = 1
	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}

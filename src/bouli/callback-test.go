package main

import "fmt"

func caller(input int) {
	func() {
		fmt.Printf("Callback funcation was called with %d input.\n", input)
	}()
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	for i := range list {
		caller(i)
	}
}

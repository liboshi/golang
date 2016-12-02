package main

import (
	"fmt"
)

func main() {
	f := closure(1)
	fmt.Println(f(1))
	fmt.Println(f(2))
	foo := closureA("Li Boshi")
	fmt.Println(foo("Hello"))
	fmt.Println(foo("Hi"))
	fmt.Println("=======")
	fmt.Println(testA())
	testC(1, 2, 3, 4, 5)

	a := func() {
		fmt.Println("Func Anonymous...")
	}
	a()
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func closureA(name string) func(string) string {
	return func(word string) string {
		return word + " " + name
	}
}

func testA() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func testB(a, b, c int) int {
	a, b, c = 1, 2, 3
	return a
}

func testC(a ...int) {
	fmt.Println(a)
	fmt.Println(a[0])
}

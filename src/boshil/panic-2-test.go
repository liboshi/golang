package main

import "fmt"

func testA() {
	fmt.Println("Func A")
}

func testB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func testC() {
	fmt.Println("Func C")
}

func main() {
	testA()
	testB()
	testC()
}

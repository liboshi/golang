package main

import (
	"fmt"
)

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1: 4] ==", p[1:4])
	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[4:] ==,", p[4:])
	q := p[:]
	append(q, 17)
	fmt.Println(q)
	fmt.Println("============================================")

	a := make([]int, 5)
	fmt.Println("a ==", a)
	b := make([]int, 0, 5)
	fmt.Println("b ==", b)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64}
var c = make(chan int, 10)

func main() {
	for i := range pow {
		fmt.Println(i)
	}

	for i, _ := range pow {
		fmt.Println(i)
	}

	for _, v := range pow {
		fmt.Println(v)
	}

	for i, v := range pow {
		fmt.Printf("2 ^ %d = %d\n", i, v)
	}

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	for v := range c {
		fmt.Println(v)
	}
}

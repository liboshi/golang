package main

import (
	"fmt"
)

func loop(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("%d ", i)
		fmt.Println("time")
		if i == 9 {
			close(ch)
		}
	}
}

func main() {
	ch := make(chan int)
	go loop(ch)
	for {
		_, open := <-ch
		if !open {
			break
		}
	}
}

package main

import "fmt"

func foo(in chan int) {
	fmt.Println(<-in)
}

// This sample will cause deadlock.
func main() {
	out := make(chan int)
	out <- 2
	go foo(out)
	// How to fix this?
	// go foo(out)
	// out<-2
}

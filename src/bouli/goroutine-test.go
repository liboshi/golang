package main

import (
	"fmt"
	"time"
)

var c chan int

func main() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout...")
	}
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	c <- 1
}

func shower(c chan int, quit chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Println(j)
		case <-quit:
			break
		}
	}
}

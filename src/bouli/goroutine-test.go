package main

import (
	"fmt"
	"time"
)

var c chan int

func main() {
	go say("world")
	say("Hello")
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	c <- 1
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
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

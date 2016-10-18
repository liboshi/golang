package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		x := <-c
		fmt.Println("received", x)
	}()

	fmt.Println("sending Boush")
	c <- "Boush"
	fmt.Println("sent Boush")
}

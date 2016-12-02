package main

import (
	"fmt"
	"time"
)

func send_data(ch chan string) {
	names := []string{"Kevin", "Frank", "Vincent", "Jacky", "Evan"}
	for _, name := range names {
		ch <- name
	}
}

func get_data(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func main() {
	ch := make(chan string)

	go send_data(ch)
	go get_data(ch)

	time.Sleep(1 * time.Second)
	fmt.Println()
}

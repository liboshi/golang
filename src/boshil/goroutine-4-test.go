package main

import (
	"fmt"
	"time"
)

func send_data(ch chan string) {
	ch <- "Kevin"
	ch <- "Frank"
	ch <- "Vincent"
	ch <- "Jacky"
	ch <- "Evan"
	close(ch)
}

func get_data(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

func main() {
	ch := make(chan string)

	go send_data(ch)
	get_data(ch)

	time.Sleep(1e9)
}

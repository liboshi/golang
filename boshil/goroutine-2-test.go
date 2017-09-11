package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	names := []string{"Kevin", "Frank", "Vincent", "Jacky", "Evan"}
	for _, name := range names {
		ch <- name
	}
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1 * time.Second)
	fmt.Println()
}

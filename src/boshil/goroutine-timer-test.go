package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e9)
	boom := time.After(10e9)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
			return
		default:
			fmt.Println("    .")
			time.Sleep(2e9)
		}
	}
}

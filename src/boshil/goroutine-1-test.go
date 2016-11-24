package main

import (
	"fmt"
	"time"
)

func long_wait() {
	fmt.Println("Begin long_wait...")
	time.Sleep(5 * time.Second)
	fmt.Println("End long_wait...")
}

func short_wait() {
	fmt.Println("Begin short_wait...")
	time.Sleep(2 * time.Second)
	fmt.Println("End short_wait...")
}

func main() {
	fmt.Println("In main...")
	go long_wait()
	go short_wait()
	fmt.Println("About to sleep in main...")
	time.Sleep(10 * time.Second)
	fmt.Println("End main...")
}

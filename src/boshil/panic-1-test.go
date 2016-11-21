package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover here")
		}
	}()
	panic("A server error happened: stopping the program")
	fmt.Println("Ending the program")
}

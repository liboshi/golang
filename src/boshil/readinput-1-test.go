package main

import "fmt"

var (
	firstName, lastName string
	i                   int
	j                   float32
	input               = "12.34,123"
	format              = "%f,%d"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hello %s %s\n", firstName, lastName)
	fmt.Sscanf(input, format, &j, &i)
	fmt.Println("From string input: ", i, j)
}

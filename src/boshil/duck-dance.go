package main

import "fmt"

type Duck interface {
	Quack()
	Walk()
}

func DuckDance(duck Duck) {
	for i := 1; i < 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct{}

func (b *Bird) Quack() {
	fmt.Println("I am quacking")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking")
}

func main() {
	b := new(Bird)
	DuckDance(b)
}

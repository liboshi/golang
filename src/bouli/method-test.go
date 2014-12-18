package main

import (
	"fmt"
)

type A struct {
	Name string
}

func main() {
	a := &A{}
	a.Print()
	a.SayHello("Boush")
}

func (a *A) Print() {
	a.Name = "Li Boshi"
	fmt.Println("A")
}

func (a *A) SayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

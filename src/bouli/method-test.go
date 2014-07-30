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
}

func (a *A) Print() {
	a.Name = "Li Boshi"
	fmt.Println("A")
}

package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type int_ int

func (a *A) Print() {
	a.Name = "Li Boshi"
	fmt.Println("A")
}

func (b *B) Print() {
	b.Name = "Li Boshi"
	fmt.Println("B")
}

func (a *A) SayHello(name string) {
	a.Name = "Boush"
	fmt.Printf("Hello %s\n", name)
}

func (a *int_) Print() {
	fmt.Println("int_")
}

func (a *int_) Increase(num int) {
	*a += int_(num)
	fmt.Println(*a)
}

type C struct {
	Name string
	Age  int
}

func (c *C) SayHello() {
	fmt.Println("Hello, I am C...")
}

func main() {
	c := C{"Boush", 29}
	c.SayHello()
}

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
	fmt.Println(a.Name)
	fmt.Println("访问控制权限由首字母大小写来控制，大写为全局，小写只能当前Package可以访问")
}

func (a *A) Print() {
	a.Name = "Li Boshi"
	fmt.Println("A")
}

func (a *A) SayHello(name string) {
	a.Name = "Boush"
	fmt.Printf("Hello %s\n", name)
}

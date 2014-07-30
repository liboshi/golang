package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := person{Name: "Li Boshi", Age: 28}
	fmt.Println(a)
}

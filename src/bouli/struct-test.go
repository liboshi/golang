package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type personA struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func testA(p *person) {
	p.Age = 29
	fmt.Println("testA", p)
}

func testB(p *person) {
	p.Age = 30
	fmt.Println("testB", *p)
}

func maina() {
	a := &person{Name: "Li Boshi", Age: 28}
	b := person{}
	b.Name = "Li Boush"
	b.Age = 28
	fmt.Println(*a)
	testA(a)
	fmt.Println(a)
	testB(a)
	fmt.Println(a)

	c := struct {
		Name string
		Age  int
	}{
		Name: "Boush",
		Age:  28}
	fmt.Println(c)

	d := &personA{
		Name: "LiBoshi",
		Age:  28,
	}
	d.Contact.Phone = "10086"
	d.Contact.City = "Beijing"
	fmt.Println(d)

	e := teacher{Name: "teacher", Age: 28, human: human{Sex: 1}}
	f := student{Name: "student", Age: 28, human: human{Sex: 0}}
	e.Sex = 2
	e.Name = "teacher1"
	e.Age = 29
	fmt.Println(e)
	fmt.Println(f)
}

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{Name: "A", B: B{Name: "B"}}
	fmt.Println(a.Name, a.B.Name)
}

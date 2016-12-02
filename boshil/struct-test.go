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

type C struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	c := new(C)
	c.Name = "Boush"
	c.Age = 29
	c.Sex = "Male"
	fmt.Println("// Using 'new' will return a pointer...")
	fmt.Println(c)

	c1 := C{Name: "Boush", Age: 29, Sex: "Male"}
	fmt.Println(c1)
}

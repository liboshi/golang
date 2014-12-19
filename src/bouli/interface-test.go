package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (p PhoneConnecter) Name() string {
	return p.name
}

func (p PhoneConnecter) Connect() {
	fmt.Println("Connect:", p.name)
}

func Disconnect(usb USB) {
	if p, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnect:", p.name)
		return
	}
}

type Boush struct {
	name string
}

func (b *Boush) Get() string {
	return b.name
}

func (b *Boush) Set(name string) {
	b.name = name
}

type Person interface {
	Get() string
	Set(string)
}

func Bind(p Person) {
	fmt.Println(p.Get())
	p.Set("Li Boush")
	fmt.Println(p.Get())
}

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
	var p Person
	p = &Boush{"Boush"}
	fmt.Println(p.Get())
	Bind(p)
}

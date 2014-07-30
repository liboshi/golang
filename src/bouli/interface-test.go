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

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}

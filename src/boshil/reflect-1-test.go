package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello everyone...")
}

func info(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("Type: ", t.Name())

	v := reflect.ValueOf(i)
	fmt.Println("Fields: ")

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%8s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < v.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%8s: %v\n", m.Name, m.Type)
	}
}

func Set(i interface{}) {
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr && v.Elem().CanSet() {
		v = v.Elem()
	} else {
		fmt.Println("Error!")
		return
	}

	switch f := v.FieldByName("Name"); f.Kind() {
	case reflect.String:
		f.SetString("Boush")
	case reflect.Int:
		f.SetInt(1)
	default:
		fmt.Println("...")
	}
}

func main() {
	/*
		m := Manager{User: User{1, "Boush", 28}, title: "Manager"}
		t := reflect.TypeOf(m)
		fmt.Printf("%#v\n", t.Field(0))
		fmt.Printf("%#v\n", t.Field(1))
		fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	*/
	u := User{1, "Name", 28}
	Set(&u)
	fmt.Println(u)
}

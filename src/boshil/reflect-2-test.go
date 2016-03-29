package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	fmt.Println("type: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())
	fmt.Println("value: ", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

	// setting a value
	fmt.Println("settablility of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settablility of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("settablility of v:", v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}

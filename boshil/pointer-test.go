package main

import "fmt"

type TestStruct struct {
	i int
	s string
}

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p
	fmt.Println(j)

	s := TestStruct{1, "A"}
	fmt.Println(&s)
}

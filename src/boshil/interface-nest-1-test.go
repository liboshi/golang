package main

import "fmt"

type GetSetter interface {
	Getter
	Setter
}

type Getter interface {
	Get() int
}

type Setter interface {
	Set(v int)
}

type Sample struct {
	value int
}

func (s *Sample) Set(v int) {
	s.value = v
}

func (s *Sample) Get() int {
	return s.value
}

func main() {
	var sam GetSetter
	s1 := new(Sample)
	sam = s1
	if _, ok := sam.(*Sample); ok {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	sam.Set(1)
	fmt.Println(sam.Get())
}

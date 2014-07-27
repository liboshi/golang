package main

import (
  "fmt"
  "time"
)

var c chan int

func ready(w string, sec int) {
  time.Sleep(time.Duration(sec) * time.Second)
  fmt.Println(w, "is ready...")
  c <- 1
}

type Boush struct {
  name string
  age  int
}

func (b *Boush) SayHello() {
  fmt.Println("Hello, This is", b.name, ". I am", b.age)
}

func main() {
  c = make(chan int)
  go ready("Tea", 2)
  go ready("Coffee", 1)
  fmt.Println("I am waiting, but not too long...")
  <- c
  <- c
  b := &Boush{"Li Bohsi", 28}
  b.SayHello()
}

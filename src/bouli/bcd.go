package main

import (
  "fmt"
  "time"
  "os"
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

func shower(c chan int, quit chan bool) {
  for {
    select {
      case j := <-c:
        fmt.Println(j)
      case <-quit:
        break
    }
  }
}

func main() {
  buf := make([]byte, 256)
  f, _ := os.Open("/etc/passwd")
  defer f.Close()
  for {
      n, _ := f.Read(buf)
    if n == 0 {
          break
      }
      os.Stdout.Write(buf[:n])
  }
}

package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	c <- 1
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
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

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit...")
			return
		}
	}
}

func fibonaccii(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	/*
				go say("world")
				say("Hello")
				a := []int{1, 2, 3, 4, 5, 6}
				c := make(chan int)
				go sum(a[:len(a)/2], c)
				go sum(a[len(a)/2:], c)
				x, y := <-c, <-c

				fmt.Println(x, y, x+y)
				c := make(chan int, 2)
				c <- 1
				c <- 2
				fmt.Println(<-c)
				c <- 3
				fmt.Println(<-c)
				fmt.Println(<-c)
			c := make(chan int)
			quit := make(chan int)
			go func() {
				for i := 0; i < 10; i++ {
					fmt.Println(<-c)
				}
				quit <- 0
			}()
			fibonacci(c, quit)
		c := make(chan int, 10)
		go fibonaccii(cap(c), c)
		for i := range c {
			fmt.Println(i)
		}
	*/

	defer fmt.Println("World")
	fmt.Println("Hello")
}

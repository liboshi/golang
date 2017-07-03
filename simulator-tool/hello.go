package main

import (
	"fmt"
	//	"math"
	//	"runtime"
	"time"
	//	"os/exec"
	"net/http"
	//	"bouli"
	"flag"
	"os"
)

var i, j int = 1, 2
var c, python, ruby bool = true, true, true

type Person struct {
	Name string
	Age  int
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type MyError struct {
	When time.Time
	What string
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Errors
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number: %f", e)
}

func run() error {
	return &MyError{
		time.Now(),
		"It doesn't work.",
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}
	return 0, nil
}

// Web server
type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// goroutine
func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

var ch chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	ch <- 1
}

// channel
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func shower(c, quit chan int) {
	for {
		select {
		case j := <-c:
			fmt.Println(j)
		case <-quit:
			fmt.Println("Quit---")
			break
		}
	}
}

// range and close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// defer: The statements after defer will be invoked before the func exit.
func foo() {
	defer fmt.Println("world")
	fmt.Println("Hello 0")
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

func fib(c, quit chan int) {
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

func test() {
	fmt.Println("Method: test()")
}

func ParseArgs() {
	if len(os.Args) <= 1 {
		fmt.Println("Arguments number is not enough...")
		os.Exit(1)
	}
	var (
		help = flag.String("h", "", "Help documentation...")
		user = flag.String("u", "Username", "User name...")
	)
	flag.Parse()
	fmt.Println("Help:", *help)
	fmt.Println("Username:", *user)
}

func main() {
	fmt.Println("Hello golang")
	fmt.Println("Hello world")
}

package main 

import (
	"fmt"
//	"math"
//	"runtime"
	"time"
//	"os"
	"net/http"
)

var i, j int = 1, 2
var c, python, ruby bool = true, true, true
type Person struct {
	Name string
	Age int
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
type Hello struct {}

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

// channel
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

// range and close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

// defer: The statements after defer will be invoked before the func exit.
func foo() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

func main() {
/*
	fmt.Printf("Hello world...\n")
	fmt.Println(math.Pi)
	fmt.Println(add(4, 5))
	a, b := swap("world", "Hello")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(i, j, c, python, ruby)
	// For statement
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	// While implementation in Golang
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
	// If...else... statement
	if sum > 128 {
		fmt.Println("=====")
	} else {
		fmt.Println("-----")
	}
	// Switch statement
	switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Linux")
		case "Windows":
			fmt.Println("Windows")
		default:
			fmt.Printf("%s\n", os)
	}
	//
	fmt.Println(time.Now().Weekday())
	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			fmt.Println("Today")
		case today + 1:
			fmt.Println("Tomorrow")
		case today + 2:
			fmt.Println("In two days")
		default:
			fmt.Println("Too far away.")
	}
	//
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning")
		case t.Hour() < 17:
			fmt.Println("Good afternoon")
		default:
			fmt.Println("Good evening")
	}
	// Stucture
	fmt.Println(Person{"Li Boshi", 28})
	p := Person{"Li Boshi", 28}
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	// Pointer
	p1 := &p
	p1.Age = 27
	fmt.Println(p)
	// Array
	var m [2]string
	m[0] = "Hello"
	m[1] = "world"
	fmt.Println(m[0], m[1])
	// Slice
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("s ==", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d ", i, s[i])
	}
	// Interface
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "\nHello world.\n")
	// Errors
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	//var h Hello
	//http.ListenAndServe("localhost:4000", h)())
*/
	// goroutine
	go say("world")
	say("Hello")
	// channel
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)
	x, y := <-c, <-c
	
	fmt.Println(x, y , x + y)
	// Buffered channel
	d := make(chan int, 2)
	fmt.Println(d)
	d <- 1
	d <- 2
	fmt.Println(<-d)
	fmt.Println(<-d)
	
	e := make(chan int, 10)
	fmt.Println(cap(e))
	go fibonacci(cap(e), e)
	for i := range e {
		fmt.Println(i)
	}

	// map
	m := make(map[string]string)
	m["name"] = "Li Boshi"
	fmt.Println(m["name"])
	foo()
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}


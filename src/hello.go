package main 

import (
	"fmt"
	"math"
	"runtime"
	"time"
	"os"
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

func run() error {
	return &MyError{
		time.Now(),
		"It doesn't work.",
	}
}

func main() {
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
}


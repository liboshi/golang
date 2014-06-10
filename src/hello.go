package main 

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

var i, j int = 1, 2
var c, python, ruby bool = true, true, true
type Person struct {
	Name string
	Age int
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
	var m [2]string
	m[0] = "Hello"
	m[1] = "world"
	fmt.Println(m[0], m[1])
}


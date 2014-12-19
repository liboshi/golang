package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work...",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v",
		float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}
	return 0, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))
}

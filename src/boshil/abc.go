package main

import (
	"fmt"
	"regexp"
)

func SayHello() {
	fmt.Println("Hello Guys.")
	fmt.Println("Hello world.")
}

func regexpSample() {
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindString)
}

package main

import (
	"fmt"
	"regexp"
)

func SayHello() {
	fmt.Println("Hello Guys.")
}

func regexpSample() {
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindString)
}

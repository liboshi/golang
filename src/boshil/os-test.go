package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	os.Mkdir("/Users/ql/test/testdir", 0777)
	fmt.Println("Create direcotry done")

	_, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("/Users/boshil/github.com/golang/README.md")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

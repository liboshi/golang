package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("/Users/ql/test/testdir", 0777)
	fmt.Println("Create direcotry done")
}

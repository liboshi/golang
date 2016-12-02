package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var execPath string
	var flagvar1 int
	duration := flag.Int("duration", 200, "Monkey test execution duration.")
	// Bind the flag to a variable.
	flag.IntVar(&flagvar1, "flagname", 1234, "help message for flagname")
	flag.StringVar(&execPath, "execPath", "/home/boshil/test", "Execution path user sepecified.")
	flag.Parse()
	fmt.Println(*duration)
	fmt.Println(flagvar1)
	fmt.Println(execPath)
	fmt.Println(os.Args[1:])
}

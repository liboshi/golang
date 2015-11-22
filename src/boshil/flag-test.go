package main

import (
	"flag"
	"fmt"
)

func main() {
	duration := flag.Int("duration", 200, "Monkey test execution duration.")
	var execPath string
	flag.StringVar(&execPath, "execPath", "/home/boshil/test", "Execution path user sepecified.")
	flag.Parse()
	fmt.Println(*duration)
	fmt.Println(execPath)
}

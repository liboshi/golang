package main

import (
  "fmt"
  "flag"
)

func main() {
  duration := flag.Int("duration", 200, "Monkey test execution duration.")
  var execPath string
  flag.StringVar(&execPath, "execPath", "/home/bouli/exec", "Execution path user sepecified.")
  flag.Parse()
  fmt.Println(*duration)
  fmt.Println(execPath)
}

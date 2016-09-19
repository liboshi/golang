package main

import (
	"boshil/cli"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	cli.Say("Boush")
	os.Exit(cli.Run(Args[1:]))
}

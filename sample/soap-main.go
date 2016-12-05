package main

import (
	"encoding/xml"
	"fmt"
	"github.com/liboshi/golang/soap"
	"os"
)

func main() {
	v := &soap.Envelope{}
	v.Header = &soap.Header{}
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("   ", "      ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

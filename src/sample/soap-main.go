package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"soap"
)

func main() {
	v := &soap.SOAPEnvelope{}
	v.Header = &soap.SOAPHeader{}
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("   ", "      ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()
}

package main

import "fmt"

func main() {
	// Initialize
	// The first method
	var m map[int]string
	// m = map[int]string{}
	m = make(map[int]string)
	fmt.Println(m)
	m[1] = "My Name: "
	m[2] = "Li Boshi"

	for _, v := range m {
		fmt.Println(v)
	}

	// The second method
	// n := map[int]string{}
	n := make(map[int]string)
	fmt.Println(n)

}

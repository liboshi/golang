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

	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)

	m3 := make(map[int]map[int]string)
	m3[1] = map[int]string{1: "a"}
	fmt.Println(m3[1][1])
}

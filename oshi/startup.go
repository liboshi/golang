package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func downloader(url, build string) {
	out, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error, Create file")
	}
	defer out.Close()
	resp, err := http.Get(url + build)
	if err != nil {
		fmt.Println("Error, No response.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error, read the response.")
	}
	fmt.Printf("%s", body)
}

func main() {
	fmt.Println("I will develop a new automation framework using Golang...")
	downloader("http://www.example.com/build?=", "100200")
}

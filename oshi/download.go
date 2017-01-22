package main

import (
	"fmt"
	. "github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	URLPREFIX    = "build web url"
	BINURLPREFIX = "build bin web url"
)

func GetResoureList(url string) []byte {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("error")
	}

	if resp.StatusCode != 200 {
		fmt.Printf("StatusCode is: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error")
	}

	return body
}

func GetLatestBuildNo(respBody []byte) int {
	js, err := NewJson(respBody)
	if err != nil {
		fmt.Println("error")
	}
	_list := js.Get("_list")
	if _list == nil {
		panic("_list is nil")
	}
	buildno, _ := _list.GetIndex(0).Get("id").Int()
	return buildno
}

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
	respBody := GetResoureList(url)
	fmt.Println(GetLatestBuildNo(respBody))
}

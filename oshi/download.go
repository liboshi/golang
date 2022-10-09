package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	. "github.com/bitly/go-simplejson"
)

const (
	URLPREFIX    = ""
	BINURLPREFIX = ""
	DLVURLPREFIX = ""
)

func URLForLatestBuild(buildType, product, branch string) string {
	url := URLPREFIX + fmt.Sprintf(`/%s/build/?product=%s`+
		`&branch=%s`+
		`&buildstate__in=succeeded,storing`+
		`&buildtype__in=beta,release`+
		`&_order_by=-id`+
		`&_limit=1`,
		buildType, product, branch)
	return url
}

func URLForDeliverable(buildno string) string {
	url := DLVURLPREFIX + buildno
	return url
}

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

func GetBuildNo(respBody []byte) int {
	js, err := NewJson(respBody)
	if err != nil {
		fmt.Println("error in GetBuildNo")
	}
	_list := js.Get("_list")
	if _list == nil {
		panic("_list is nil")
	}
	buildno, _ := _list.GetIndex(0).Get("id").Int()
	return buildno
}

func GetLatestBuildNo(respBody []byte) int {
	return GetBuildNo(respBody)
}

func GetDeliverableUrl(respBody []byte) string {
	js, err := NewJson(respBody)
	if err != nil {
		fmt.Println("error in GetDeliverableUrl")
	}
	_list := js.Get("_list")
	if _list == nil {
		panic("_list is nil")
	}
	deliverableUrl, _ := _list.GetIndex(0).Get("_deliverables_url").String()
	return deliverableUrl
}

func GetBinaryPath(respBody []byte) string {
	var binaryPath string
	js, err := NewJson(respBody)
	if err != nil {
		fmt.Println("error in GetBinaryPath")
	}
	_list := js.Get("_list").MustArray()
	if _list == nil {
		panic("_list is nil")
	}
	for _, d := range _list {
		binaryPath = d.MustMap()["path"]
		matched, _ := regexp.MatchString("*VMware-viewagent*", binaryPath)
		if matched {
			return binaryPath
		}
	}
	return ""
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
	/*
		buildType := "ob"
		product := "viewclientwin"
		branch := "crt-17q2"
		url := URLForLatestBuild(buildType, product, branch)
		respBody := GetResoureList(url)
		fmt.Println(GetLatestBuildNo(respBody))
		deliverableUrl := GetDeliverableUrl(respBody)
		url = URLPREFIX + deliverableUrl
		fmt.Println(url)
		respBody = GetResoureList(url)
		fmt.Println(GetBinaryPath(respBody))
	*/
	url := URLForDeliverable("5612119")
	fmt.Println(url)
	respBody := GetResoureList(url)
	fmt.Println(GetBinaryPath(respBody))
}

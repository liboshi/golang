package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func upgradeGolang(goSrcFile string) {
	cmd := exec.Command("which", "go")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := out.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Golang is installed...")
		fmt.Println("Upgrade Golang...")
	}
	cmd = exec.Command("sudo", "rm", "-rf", "/usr/local/go")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Remove Golang failed...")
	}
	cmd = exec.Command("sudo", "tar", "-C", "/usr/local", "zxvf", goSrcFile)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Install Golang failed...")
	}
}

func fetchGolangInstaller(version string) {
	goDownloadUrl := "https://storage.googleapis.com/golang/go1.4.darwin-amd64-osx10.8.tar.gz"
	resp, err := http.Get(goDownloadUrl)
	if err != nil {
		fmt.Println("Download golang installer failed...")
	}
	// Close the handler
	defer resp.Body.close()
	// Download...
	n, err := io.Copy(out, resp.Body)
}

func main() {
	argNum := len(os.Args)
	goSrcFile := "~/Downloads/go1.4rc2.linux-amd64.tar.gz"
	if argNum != 2 {
		fmt.Println("The number of arguments is not correct...")
		fmt.Println("Default version of Golang will be installed...")
	}
	upgradeGolang(goSrcFile)
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func brew_install(pkg string, args ...int) {
	cmd := exec.Command("brew", "list", pkg, "--versions")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	} else {
		output, err := out.ReadString('\n')
		if err != nil {
			cmd := exec.Command("brew", "install", pkg)
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println(output)
	}
}

func install_github_bundle(user, pkg string) {
	dir := "~/.vim/bundle/" + pkg
	fmt.Println(dir)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Found the target directory...")
		}
	} else {
		cmd := exec.Command("git", "clone",
			"https://github.com/"+user+pkg+" ~/.vim/bundle/"+pkg)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	install_github_bundle("boush", "jedi-vim")
}

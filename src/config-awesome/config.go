package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func brewInstall(pkg string, args ...int) {
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

func installGithubBundle(user, pkg string) {
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

func updateOrInstallBrew() {
	fmt.Println("Update or install brew")
	cmd := exec.Command("which", "brew > /dev/null", "||", "ruby -e", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	installGithubBundle("boush", "jedi-vim")
}

package main

import (
	"os"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func check(e error, message string) {
	if e != nil {
		Error.Println(e, message)
	}
}

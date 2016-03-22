package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "/Users/boshil/works/tmp/cpuprifle",
	"Write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		time.Sleep(30e9)

		pprof.StopCPUProfile()
		fmt.Println("Stop profilling after 30 seconds")
		defer f.Close()
	}
}

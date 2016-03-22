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

func loop() {
	for i := 0; i < 1000; i++ {
		time.Sleep(1e7)
		fmt.Println(".....")
	}
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)

		loop()
		f.Close()
		defer pprof.StopCPUProfile()
	}
}

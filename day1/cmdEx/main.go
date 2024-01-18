package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/Fabulous-Fadz/aoc23/day1/trebuchet"
	"github.com/Fabulous-Fadz/aoc23/internal"
)

var (
	profile  = flag.Bool("profile", false, "Should we run a profiler on this?")
	fileData = internal.Must(os.ReadFile("../input.txt"))
)

func main() {
	flag.Parse()
	if *profile {
		cpuProfile := internal.Must(os.Create("cpuprofile"))
		memProfile := internal.Must(os.Create("memprofile"))
		if err := pprof.StartCPUProfile(cpuProfile); err == nil {
			defer pprof.StopCPUProfile()
		}
		defer pprof.WriteHeapProfile(memProfile)
	}

	items := [][]byte{}

	foo := strings.Split(string(fileData), "\n")
	for _, bar := range foo {
		items = append(items, []byte(bar))
	}

	total := trebuchet.DecodeAllWords(items)
	fmt.Printf("Total with words: %v\n", total)
}

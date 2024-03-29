package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/Fabulous-Fadz/aoc23/day1/trebuchet"
	"github.com/Fabulous-Fadz/aoc23/internal"
)

var (
	fileData   = internal.Must(os.ReadFile("../input.txt"))
	cpuProfile = internal.Must(os.Create("cpuprofile"))
	memProfile = internal.Must(os.Create("memprofile"))
)

func main() {
	if err := pprof.StartCPUProfile(cpuProfile); err == nil {
		defer pprof.StopCPUProfile()
	}
	defer pprof.WriteHeapProfile(memProfile)

	items := [][]byte{}

	foo := strings.Split(string(fileData), "\n")
	for _, bar := range foo {
		items = append(items, []byte(bar))
	}

	total := trebuchet.DecodeAll(items)
	fmt.Println(total)
}

func must[T any](result T, err error) T {
	if err != nil {
		log.Fatal("Could not complete operation.")
	}
	return result
}

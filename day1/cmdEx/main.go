package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Fabulous-Fadz/aoc23/day1/trebuchet"
	"github.com/Fabulous-Fadz/aoc23/internal"
)

var fileData = internal.Must(os.ReadFile("../input.txt"))

func main() {
	items := [][]byte{}

	foo := strings.Split(string(fileData), "\n")
	for _, bar := range foo {
		items = append(items, []byte(bar))
	}

	total := trebuchet.DecodeAllWords(items)
	fmt.Printf("Total with words: %v\n", total)
}

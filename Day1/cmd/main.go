package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	trebuchet "github.com/Fabulous-Fadz/aoc23/Day1/Trebuchet"
)

var fileData = must(os.ReadFile("input.txt"))

func main() {
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

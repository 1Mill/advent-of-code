package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf []int
type Elves []Elf

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchElves(file string) Elves {
	// * Import file as bytes
	data, err := os.ReadFile(file)
	check(err)

	// * Turn bytes into string
	contnet := string(data)

	// * Split string by double new lines to get elf sub-strings
	// * Split elf sub-strings by new lines
	// * Convert lines to numbers
	groups := strings.Split(contnet, "\n\n")
	elves := make(Elves, len(groups))

	for i, group := range groups {
		lines := strings.Split(group, "\n")
		elf := make(Elf, len(lines))

		for j, line := range lines {
			if line == "" {
				continue
			}

			num, err := strconv.Atoi(line)
			check(err)

			elf[j] = num
		}

		elves[i] = elf
	}

	return elves
}

func main() {
	file := "./example.txt"
	// file := "./input.txt"

	elves := fetchElves(file)

	fmt.Println("\n")
	fmt.Println(elves)
}

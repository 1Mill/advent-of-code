package main

import (
	"fmt"
	"os"
	"strings"
)

type Calorie int
type Elf []Calorie
type Elves []Elf

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file := "./example.txt"
	// file := "./input.txt"

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

	// for _, group := range groups {
	// for j, line := range strings.Split(group, "\n") {
	// 	elf := make([]int, len())

	// 	if line == "" {
	// 		continue
	// 	}

	// 	num, err := strconv.Atoi(line)
	// 	check(err)

	// 	elves[i][j] = num
	// }
	// }

	fmt.Println(elves)
}

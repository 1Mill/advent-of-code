package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf []int
type Elves []Elf

func addTop(tops []int, i, num int) []int {
	// * Insert empty value at index
	// * [1, 2, 3, ____, 4, 5, 6]
	newTops := append(tops[:i+1], tops[i:]...)

	// * Assign value at index to given number
	// * [1, 2, 3, 8080, 4, 5, 6]
	newTops[i] = num

	// * Return array with original length, dropping the last item
	// * [1, 2, 3, 8080, 4, 5]
	return newTops[:len(tops)]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func computeMax(elves Elves) int {
	maxCalories := 0

	for _, elf := range elves {
		sum := sumElf(elf)

		if maxCalories < sum {
			maxCalories = sum
		}
	}

	return maxCalories
}

func fetchElves(file string) Elves {
	// * Import file as bytes
	data, err := os.ReadFile(file)
	check(err)

	// * Turn bytes into string
	contnet := string(data)

	// * Split string by double new lines to get elf sub-strings
	groups := strings.Split(contnet, "\n\n")
	elves := make(Elves, len(groups))

	for i, group := range groups {
		// * Split content group by new lines to get elf calories
		lines := strings.Split(group, "\n")
		elf := make(Elf, len(lines))

		for j, line := range lines {
			if line == "" {
				continue
			}

			// * Convert string to number
			num, err := strconv.Atoi(line)
			check(err)

			elf[j] = num
		}

		elves[i] = elf
	}

	return elves
}

func sumElf(elf Elf) int {
	sum := 0
	for _, calories := range elf {
		sum += calories
	}
	return sum
}

func topNMax(elves Elves, n int) int {
	tops := make([]int, n)

	for _, elf := range elves {
		num := sumElf(elf)
		for j, val := range tops {
			if num > val {
				tops = addTop(tops, j, num)
				break
			}
		}
	}

	sum := 0
	for _, v := range tops {
		sum += v
	}

	return sum
}

func main() {
	// file := "./example.txt"
	file := "./input.txt"

	elves := fetchElves(file)

	max := computeMax(elves)
	fmt.Println("Part 1, Max:", max)

	n := 3
	topNMax := topNMax(elves, n)
	fmt.Printf("Part 2, Top %d Sum: %d\n", n, topNMax)
}

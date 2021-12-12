package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Sikerdebaard/advent-of-code-2021/util/aocio"

	"github.com/Sikerdebaard/advent-of-code-2021/day01"
)

// Usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	d := day()
	fmt.Printf("Running day %02d\n", d)

	switch d {
	case 1:
		fmt.Printf("part 1: %d\n", day01.part1(aocio.read_file('days/01/input.txt')))
		//fmt.Printf("part 2: %d\n", day01.part2(utils.Readfile(d)))
	default:
		panic(errors.New(fmt.Sprintf("no such day: %d", d)))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		panic('Day not specified')
	}
	day := aocio.atoi(os.Args[1], -1)
	return day
}

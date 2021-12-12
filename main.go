package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Sikerdebaard/advent-of-code-2021/utils/aocio"

	"github.com/Sikerdebaard/advent-of-code-2021/days/day01"
	"github.com/Sikerdebaard/advent-of-code-2021/days/day02"
)

// Usage: go run main.go <NN>
// assumes input is in day<NN>/input.txt
func main() {
	d := day()
	fmt.Printf("Running day %02d\n", d)

	switch d {
	case 1:
		fmt.Printf("part 1: %s\n", day01.Part1(aocio.ReadFile("days/day01/input.txt")))
		fmt.Printf("part 2: %s\n", day01.Part2(aocio.ReadFile("days/day01/input.txt")))
	case 2:
		fmt.Printf("part 1: %s\n", day02.Part1(aocio.ReadFile("days/day02/input.txt")))
		fmt.Printf("part 2: %s\n", day02.Part2(aocio.ReadFile("days/day02/input.txt")))
	default:
		panic(errors.New(fmt.Sprintf("no such day: %d", d)))
	}
}

// Reads day from os.Args.
func day() int {
	if len(os.Args) == 1 {
		panic("Day not specified")
	}
	day := aocio.Atoi(os.Args[1], -1)
	return day
}

package day02

import (
	"fmt"
	"strings"
	"strconv"
)


func Part1(input string) string {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	depth := 0
	horz := 0

	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], " ")
		amount, _ := strconv.Atoi(tokens[1])

		if (tokens[0] == "forward") {
			horz += amount
		} else if (tokens[0] == "up") {
			depth -= amount
		} else if (tokens[0] == "down") {
			depth += amount
		}
	}

	ret := fmt.Sprintf("depth: %d horizontal: %d, result: %d", depth, horz, depth*horz)

	return ret
}

func Part2(input string) string {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	depth := 0
	horz := 0
	aim := 0

	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], " ")
		amount, _ := strconv.Atoi(tokens[1])

		if (tokens[0] == "forward") {
			horz += amount
			depth += aim * amount
		} else if (tokens[0] == "up") {
			aim -= amount
		} else if (tokens[0] == "down") {
			aim += amount
		}
	}

	ret := fmt.Sprintf("depth: %d horizontal: %d, aim: %d, result: %d", depth, horz, aim, depth*horz)

	return ret
}


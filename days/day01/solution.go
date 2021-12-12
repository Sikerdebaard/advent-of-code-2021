package day01

import (
	"fmt"
	"github.com/Sikerdebaard/advent-of-code-2021/utils/aocio"
)

func Part1(input string) string {
	var values = []int{}

	values = aocio.InputToIntArr(input)

	inc := 0
	dec := 0
	prev := values[0]

	for i := 1; i < len(values); i++ {
		if (values[i] > prev) {
			inc++
		} else {
			dec++
		}
		prev = values[i]
	}

	ret := fmt.Sprintf("increase: %d, decrease: %d", inc, dec)

	return ret
}

func Part2(input string) string {
	var values = []int{}

	values = aocio.InputToIntArr(input)

	inc := 0
	dec := 0

	for i := 3; i < len(values); i++ {
		slide_a := values[i-3] + values[i-2] + values[i-1]
		slide_b := values[i-2] + values[i-1] + values[i]
		if (slide_b > slide_a) {
			inc++
		} else {
			dec++
		}
	}

	ret := fmt.Sprintf("increase: %d, decrease: %d", inc, dec)

	return ret
}

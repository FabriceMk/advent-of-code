package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input
var input string

//go:embed input-test
var test_input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input file")
	}

	test_input = strings.TrimRight(test_input, "\n")
	if len(test_input) == 0 {
		panic("empty input-test file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var res string

	if part == 1 {
		resPart1 := part1(input)
		res = strconv.Itoa(resPart1)

	} else {
		resPart2 := part2(input)
		res = strconv.Itoa(resPart2)
	}

	fmt.Println("Output:")
	fmt.Println(res)
}

func part1(input string) int {
	var sum int
	for _, item := range strings.Split(input, "\n") {
		var firstDigit rune
		var lastDigit rune

		for _, character := range item {
			if unicode.IsDigit(character) {
				if firstDigit == 0 {
					firstDigit = character
					lastDigit = character
				} else {
					lastDigit = character
				}
			}
		}

		number, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		if err == nil {
			sum += number
		}
	}
	return sum
}

func part2(input string) int {
	return 2
}

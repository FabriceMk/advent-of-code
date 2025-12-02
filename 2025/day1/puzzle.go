package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input-full
var input string

//go:embed input-test
var test_input string

//go:embed input-test-2
var test_input2 string

func part1(input string) int {
	answer := 0

	position := 50

	for _, row := range strings.Split(input, "\n") {
		var rotationDigit = string(row[0])
		var distanceDigit, _ = strconv.Atoi(string(row[1:]))

		var remain = distanceDigit % 100

		if rotationDigit == "L" {
			position = (100 + (position - remain)) % 100
		}

		if rotationDigit == "R" {
			position = (position + distanceDigit) % 100
		}

		if position == 0 {
			answer++
		}
	}

	return answer
}

func part2(input string) int {
	answer := 0

	position := 50

	for _, row := range strings.Split(input, "\n") {
		var rotationDigit = string(row[0])
		var distanceDigit, _ = strconv.Atoi(string(row[1:]))

		fullCircles := distanceDigit / 100

		clicksThroughZero := fullCircles

		var remain = distanceDigit % 100

		if rotationDigit == "L" {
			if position-remain <= 0 {
				if position != 0 {
					clicksThroughZero++
				}

				position = (100 + (position - remain)) % 100
			} else {
				position = (position - remain)
			}
		}

		if rotationDigit == "R" {
			if position+remain >= 100 {
				clicksThroughZero++
			}

			position = (position + remain) % 100
		}

		answer += clicksThroughZero
	}

	return answer
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input file")
	}

	test_input = strings.TrimRight(test_input, "\n")
	if len(test_input) == 0 {
		panic("empty input-test file")
	}

	test_input2 = strings.TrimRight(test_input2, "\n")
	if len(test_input2) == 0 {
		panic("empty input-test-2 file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("##########")
	fmt.Println("Running part", part)

	var res string

	start := time.Now()

	if part == 1 {
		resPart1 := part1(input)
		res = strconv.Itoa(resPart1)

	} else {
		resPart2 := part2(input)
		res = strconv.Itoa(resPart2)
	}

	timeElapsed := time.Since(start)

	fmt.Println("Execution time: ", timeElapsed)
	fmt.Println("Output:")
	fmt.Println(res)
	fmt.Println("##########")
}

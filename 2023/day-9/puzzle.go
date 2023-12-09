package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input
var input string

//go:embed input-test
var test_input string

//go:embed input-test-2
var test_input2 string

func part1(input string) int {
	sum := 0

	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			continue
		}

		lineStr := strings.Fields(line)
		numbers := []int{}
		for _, character := range lineStr {
			charAsInt, _ := strconv.Atoi(character)
			numbers = append(numbers, charAsInt)
		}

		pyramid := [][]int{}

		pyramid = append(pyramid, numbers)

		loops := 0
		reachedStability := false

		for !reachedStability {
			reachedStability = true
			calcDiff := []int{}

			for i := 0; i < len(pyramid[loops])-1; i++ {
				diff := pyramid[loops][i+1] - pyramid[loops][i]

				if diff != 0 {
					reachedStability = false
				}

				calcDiff = append(calcDiff, diff)
			}

			pyramid = append(pyramid, calcDiff)

			loops++
		}

		intermediateResults := make([]int, loops+1)

		for i := loops; i > 0; i-- {
			intermediateResult := pyramid[i-1][len(pyramid[i-1])-1] + intermediateResults[i]

			intermediateResults[i-1] = intermediateResult
		}

		sum += intermediateResults[0]
	}

	return sum
}

func part2(input string) int {
	sum := 0

	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			continue
		}

		lineStr := strings.Fields(line)
		numbers := []int{}
		for _, character := range lineStr {
			charAsInt, _ := strconv.Atoi(character)
			numbers = append(numbers, charAsInt)
		}

		pyramid := [][]int{}

		pyramid = append(pyramid, numbers)

		loops := 0
		reachedStability := false

		for !reachedStability {
			reachedStability = true
			calcDiff := []int{}

			for i := 0; i < len(pyramid[loops])-1; i++ {
				diff := pyramid[loops][i+1] - pyramid[loops][i]

				if diff != 0 {
					reachedStability = false
				}

				calcDiff = append(calcDiff, diff)
			}

			pyramid = append(pyramid, calcDiff)

			loops++
		}

		intermediateResults := make([]int, loops+1)

		for i := loops; i > 0; i-- {
			intermediateResult := pyramid[i-1][0] - intermediateResults[i]

			intermediateResults[i-1] = intermediateResult
		}

		sum += intermediateResults[0]
	}

	return sum
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

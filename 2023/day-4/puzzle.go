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

func part1(input string) int {
	var sum int

	inputLines := strings.Split(input, "\n")
	if len(inputLines) == 0 {
		panic("malformed input")
	}

	for _, card := range inputLines {
		if len(card) == 0 {
			continue
		}

		wonCards := getWinningNumbersCount(card)

		cardSum := 0
		for i := 0; i < wonCards; i++ {
			if cardSum == 0 {
				cardSum = 1
			} else {
				cardSum *= 2
			}
		}

		sum += cardSum
	}

	return sum
}

func getWinningNumbersCount(card string) int {
	dataSection := strings.Split(card, ":")[1]
	numbersSection := strings.Split(dataSection, "|")

	winningNumbers := strings.Fields(numbersSection[0])
	drawnNumbers := strings.Fields(numbersSection[1])

	winningNumbersMap := map[int]bool{}
	for _, winningNumberStr := range winningNumbers {
		winningNumber, _ := strconv.Atoi(winningNumberStr)
		winningNumbersMap[winningNumber] = true
	}

	won := 0

	for _, drawnNumberStr := range drawnNumbers {
		drawnNumber, _ := strconv.Atoi(drawnNumberStr)
		if winningNumbersMap[drawnNumber] {
			won++
		}
	}

	return won
}

func part2(input string) int {
	inputLines := strings.Split(input, "\n")
	if len(inputLines) == 0 {
		panic("malformed input")
	}

	sum := 0

	computedBranches := map[int]int{}

	for i := 0; i < len(inputLines); i++ {
		sum += calculateCardsCount(inputLines, i, computedBranches)
	}

	return sum
}

func calculateCardsCount(inputLines []string, startingIndex int, computedBranches map[int]int) int {
	virtualRemainingLength := len(inputLines) - startingIndex

	if virtualRemainingLength <= 0 {
		return 0
	}

	if virtualRemainingLength == 1 {
		return 1
	}

	// The usage of this lookup table was done after initial implementation. Went from 5.7s to 580ns on a M1 Pro
	if computedBranches[startingIndex] != 0 {
		return computedBranches[startingIndex]
	}

	wonCards := getWinningNumbersCount(inputLines[startingIndex])

	count := 0

	for i := 0; i < wonCards; i++ {
		count += calculateCardsCount(inputLines, startingIndex+i+1, computedBranches)
	}

	if computedBranches[startingIndex] == 0 {
		computedBranches[startingIndex] = count + 1
	}

	return count + 1
}

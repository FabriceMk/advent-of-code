package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//go:embed input-full
var input string

//go:embed input-test
var test_input string

//go:embed input-test-2
var test_input2 string

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

func part1(input string) int {
	var sum int

	bagContent := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// Approach: I went for a parser solution that is not really generic as it relies heavily on the input to be correctly formatted.
	// So it's specific but I wanted to try to have a speedy approach.

	for index, item := range strings.Split(input, "\n") {
		gameID := index + 1

		if len(item) == 0 {
			break
		}

		// Skip the 'Game {ID}:' substring as ID can be deducted
		semicolonIndex := strings.Index(item, ":")
		if semicolonIndex == -1 {
			// A line is malformed so it is skipped (but it's not supposed to happen)
			continue
		}

		gameRunes := []rune(item[semicolonIndex+1:])

		currentBallCounter := 0
		var currentBallColor strings.Builder
		currentGameSet := map[string]int{
			RED:   0,
			GREEN: 0,
			BLUE:  0,
		}

		gameRunesCurrentIndex := 0
		earlyBreak := false

		for gameRunesCurrentIndex < len(gameRunes) {
			currentRune := gameRunes[gameRunesCurrentIndex]

			if currentRune == ' ' {
				gameRunesCurrentIndex++
				continue
			}

			if currentRune == ';' {
				commitGameSet(&currentBallCounter, &currentBallColor, currentGameSet)

				if !isValidGameSet(currentGameSet, bagContent) {
					// if the current game set is invalid, we don't need to process the rest of this game
					earlyBreak = true
					break
				}
			} else if currentRune == ',' {
				commitGameSet(&currentBallCounter, &currentBallColor, currentGameSet)
			} else if unicode.IsDigit(currentRune) {
				currentDigit, _ := strconv.Atoi(string(currentRune))
				currentBallCounter = currentBallCounter*10 + currentDigit
			} else {
				currentBallColor.WriteRune(currentRune)

				if gameRunesCurrentIndex == len(gameRunes)-1 {
					commitGameSet(&currentBallCounter, &currentBallColor, currentGameSet)
				}
			}

			gameRunesCurrentIndex++
		}

		if !earlyBreak && isValidGameSet(currentGameSet, bagContent) {
			sum += gameID
		}
	}

	return sum
}

func isValidGameSet(gameSet map[string]int, bagContent map[string]int) bool {
	return gameSet[RED] <= bagContent[RED] &&
		gameSet[GREEN] <= bagContent[GREEN] &&
		gameSet[BLUE] <= bagContent[BLUE]
}

func commitGameSet(currentBallCounter *int, currentBallColor *strings.Builder, currentGameSet map[string]int) {
	currentGameSet[currentBallColor.String()] = *currentBallCounter
	*currentBallCounter = 0
	currentBallColor.Reset()
}

func part2(input string) int {
	var sum int

	for _, item := range strings.Split(input, "\n") {
		if len(item) == 0 {
			break
		}

		// Skip the 'Game {ID}:' substring as ID can be deducted
		semicolonIndex := strings.Index(item, ":")
		if semicolonIndex == -1 {
			// A line is malformed so it is skipped (but it's not supposed to happen)
			continue
		}

		gameRunes := []rune(item[semicolonIndex+1:])

		currentBallCounter := 0
		var currentBallColor strings.Builder
		minimalNeededBallsCount := map[string]int{
			RED:   0,
			GREEN: 0,
			BLUE:  0,
		}

		gameRunesCurrentIndex := 0

		for gameRunesCurrentIndex < len(gameRunes) {
			currentRune := gameRunes[gameRunesCurrentIndex]

			if currentRune == ' ' {
				gameRunesCurrentIndex++
				continue
			}

			if currentRune == ';' {
				compareMinimalBallsCount(&currentBallCounter, &currentBallColor, minimalNeededBallsCount)
			} else if currentRune == ',' {
				compareMinimalBallsCount(&currentBallCounter, &currentBallColor, minimalNeededBallsCount)
			} else if unicode.IsDigit(currentRune) {
				currentDigit, _ := strconv.Atoi(string(currentRune))

				if currentBallCounter <= 0 {
					currentBallCounter = currentDigit
				} else {
					currentBallCounter = currentBallCounter*10 + currentDigit
				}
			} else {
				currentBallColor.WriteRune(currentRune)

				if gameRunesCurrentIndex == len(gameRunes)-1 {
					compareMinimalBallsCount(&currentBallCounter, &currentBallColor, minimalNeededBallsCount)
				}
			}

			gameRunesCurrentIndex++
		}

		sum += minimalNeededBallsCount[RED] * minimalNeededBallsCount[GREEN] * minimalNeededBallsCount[BLUE]
	}

	return sum
}

func compareMinimalBallsCount(currentBallCounter *int, currentBallColor *strings.Builder, minimalNeededBallsCount map[string]int) {
	if minimalNeededBallsCount[currentBallColor.String()] < *currentBallCounter {
		minimalNeededBallsCount[currentBallColor.String()] = *currentBallCounter
	}

	*currentBallCounter = 0
	currentBallColor.Reset()
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

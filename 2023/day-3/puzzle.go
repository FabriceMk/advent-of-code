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

//go:embed input
var input string

//go:embed input-test
var test_input string

//go:embed input-test-2
var test_input2 string

func part1(input string) int {
	var sum int

	// Approach: I didn't want dark magic but I couldn't resist to go with a small trick to not have to
	// check bounds.
	// The chosen algorithm is quite straightforward but probably not the most efficient.

	inputLines := strings.Split(input, "\n")
	if len(inputLines) == 0 {
		panic("malformed input")
	}

	originalLineLength := len(inputLines[0])

	// Prepare a padded matrix of runes that will contain the full schematics to help with the traversal.
	schematics := preparePaddedMatrix(&inputLines, originalLineLength)

	for i := 0; i < len(inputLines)+2; i++ {
		currentScannedNumber := 0

		for j := 0; j < originalLineLength+2; j++ {
			currentRune := schematics[i][j]

			if unicode.IsDigit(currentRune) {
				currentRuneValue, _ := strconv.Atoi(string(currentRune))
				currentScannedNumber = currentScannedNumber*10 + currentRuneValue
			} else {
				if currentScannedNumber > 0 {
					numberLength := len(strconv.Itoa(currentScannedNumber))
					foundSymbol := false

					for verticalIndex := i - 1; verticalIndex <= i+1; verticalIndex++ {
						for horizontalIndex := j - numberLength - 1; horizontalIndex <= j; horizontalIndex++ {
							scannedCharacter := schematics[verticalIndex][horizontalIndex]
							if !unicode.IsDigit(scannedCharacter) && scannedCharacter != '.' {
								foundSymbol = true
							}
						}
					}

					if foundSymbol {
						sum += currentScannedNumber
					}

					currentScannedNumber = 0
				}
			}
		}
	}

	return sum
}

func preparePaddedMatrix(inputLines *[]string, originalLineLength int) [][]rune {
	// Prepare a matrix of runes that will contain the full schematic
	var matrix [][]rune

	firstDummyLine := make([]rune, originalLineLength+2)
	for i := range firstDummyLine {
		firstDummyLine[i] = '.'
	}
	matrix = append(matrix, firstDummyLine)

	for _, item := range *inputLines {
		if len(item) == 0 {
			break
		}

		line := make([]rune, originalLineLength+2)
		line[0] = '.'
		for index, character := range item {
			line[index+1] = character
		}
		line[originalLineLength+1] = '.'

		matrix = append(matrix, line)
	}

	var lastDummyLine = make([]rune, originalLineLength+2)
	copy(lastDummyLine, firstDummyLine)

	matrix = append(matrix, lastDummyLine)

	return matrix
}

func part2(input string) int {
	var sum int

	// Approach: I didn't want dark magic but I couldn't resist to go with a small trick to not have to
	// check bounds.
	// The chosen algorithm is quite straightforward but probably not the most efficient.

	inputLines := strings.Split(input, "\n")
	if len(inputLines) == 0 {
		panic("malformed input")
	}

	originalLineLength := len(inputLines[0])

	// Prepare a padded matrix of runes that will contain the full schematics to help with the traversal.
	schematics := preparePaddedMatrix(&inputLines, originalLineLength)
	paddedLineLength := originalLineLength + 2

	// Each '*' found with its attached numbers. Each gear is identified by its index on a
	// virtual single dimension array representing the matrix.
	possibleGearList := map[int][]int{}

	for i := 0; i < len(inputLines)+2; i++ {
		currentScannedNumber := 0

		for j := 0; j < originalLineLength+2; j++ {
			currentRune := schematics[i][j]

			if unicode.IsDigit(currentRune) {
				currentRuneValue, _ := strconv.Atoi(string(currentRune))
				currentScannedNumber = currentScannedNumber*10 + currentRuneValue
			} else {
				if currentScannedNumber > 0 {
					numberLength := len(strconv.Itoa(currentScannedNumber))

					for verticalIndex := i - 1; verticalIndex <= i+1; verticalIndex++ {
						for horizontalIndex := j - numberLength - 1; horizontalIndex <= j; horizontalIndex++ {
							scannedCharacter := schematics[verticalIndex][horizontalIndex]
							if scannedCharacter == '*' {
								potentialGearIndex := verticalIndex*paddedLineLength + horizontalIndex
								possibleGearList[potentialGearIndex] = append(possibleGearList[potentialGearIndex], currentScannedNumber)
							}
						}
					}

					currentScannedNumber = 0
				}
			}
		}
	}

	for _, attachedNumbers := range possibleGearList {
		if len(attachedNumbers) == 2 {
			sum += attachedNumbers[0] * attachedNumbers[1]
		}
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

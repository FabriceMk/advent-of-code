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
				}

				lastDigit = character
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
	var sum int

	numberMapping := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for _, item := range strings.Split(input, "\n") {
		itemRunes := []rune(item)

		var firstDigit rune
		var lastDigit rune

		var index = 0

		for index < len(item) {
			currentRune := itemRunes[index]

			if unicode.IsDigit(currentRune) {
				if firstDigit == 0 {
					firstDigit = currentRune
				}

				lastDigit = currentRune

				index++
				continue
			}

			for digitLetter, digitValue := range numberMapping {
				foundDigitLetterAt := strings.Index(item[index:], digitLetter)

				if foundDigitLetterAt == 0 {
					if firstDigit == 0 {
						firstDigit = digitValue
					}

					lastDigit = digitValue
					break
				}
			}

			index++
		}

		number, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
		if err == nil {
			sum += number
		}
	}

	return sum
}

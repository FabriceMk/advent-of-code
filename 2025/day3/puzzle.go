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

	for _, bank := range strings.Split(input, "\n") {
		highestJolt := 0
		var firstBatteryHighest int
		var secondBatteryHighest int

		firstCurrent, _ := strconv.Atoi(bank[0:1])
		secondCurrent, _ := strconv.Atoi(bank[1:2])

		for _, currentBattery := range bank {
			if firstCurrent == 0 {
				firstCurrent = currentBattery
				continue
			}

			if secondCurrent == 0 {
				secondCurrent = currentBattery
				continue
			}

			currentJolt, _ := strconv.Atoi(fmt.Sprintf("%b", firstCurrent) + fmt.Sprintf("%b", secondCurrent))

			if currentJolt > highestJolt {
				highestJolt = currentJolt
				firstBatteryHighest = firstCurrent
				secondBatteryHighest = secondCurrent
			}
		}

		answer += highestJolt
	}

	return answer
}

func part2(input string) int {
	answer := 0

	for _, idRangeStr := range strings.Split(input, ",") {
		idRange := strings.Split(idRangeStr, "-")
		rangeStart, _ := strconv.Atoi(idRange[0])
		rangeEnd, _ := strconv.Atoi(idRange[1])

		for currentSequence := rangeStart; currentSequence <= rangeEnd; currentSequence++ {
			strRepresentation := strconv.Itoa(currentSequence)
			strRepresentationLength := len(strRepresentation)

			for j := 1; j <= strRepresentationLength/2; j++ {
				pattern := strRepresentation[0:j]
				patternLength := len(pattern)

				if strRepresentationLength%patternLength != 0 {
					continue
				}

				earlyBreak := false

				for k := patternLength; k <= strRepresentationLength-patternLength; k += patternLength {
					subStr := strRepresentation[k : k+patternLength]
					if subStr != pattern {
						earlyBreak = true
						break
					}
				}

				if !earlyBreak {
					answer += currentSequence
					break
				}
			}
		}
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

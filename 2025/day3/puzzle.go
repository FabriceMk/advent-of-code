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

		var firstCurrent int

		for firstBatteryIndex := 0; firstBatteryIndex < len(bank)-1; firstBatteryIndex++ {
			currentFirstBattery := string(bank[firstBatteryIndex])

			firstCurrent, _ = strconv.Atoi(currentFirstBattery)

			if firstCurrent <= firstBatteryHighest {
				continue
			}

			firstBatteryHighest = firstCurrent

			for secondBatteryIndex := firstBatteryIndex + 1; secondBatteryIndex < len(bank); secondBatteryIndex++ {
				currentSecondBattery := string(bank[secondBatteryIndex])

				secondCurrent, _ := strconv.Atoi(currentSecondBattery)

				currentJolt := firstCurrent*10 + secondCurrent

				if currentJolt > highestJolt {
					highestJolt = currentJolt
				}
			}
		}

		answer += highestJolt
	}

	return answer
}

func part2(input string) int {
	answer := 0

	for _, bank := range strings.Split(input, "\n") {
		answer += checkHighest(bank, 12)
	}

	return answer
}

func checkHighest(input string, choicesNumber int) int {

	if len(input) <= 2 {
		result, _ := strconv.Atoi(input)
		return result
	}

	return 0
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

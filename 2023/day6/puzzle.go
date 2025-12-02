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

type Mapping struct {
	mapsTo   string
	mappings [][]int
}

func part1(input string) int {
	inputLines := strings.Split(input, "\n")

	times := strings.Fields(strings.Split(inputLines[0], ":")[1])
	distances := strings.Fields(strings.Split(inputLines[1], ":")[1])

	victories := []int{}

	for i := 0; i < len(times); i++ {
		timeValue, _ := strconv.Atoi(times[i])
		distanceValue, _ := strconv.Atoi(distances[i])
		possibleWins := 0

		for pressDuration := 1; pressDuration < timeValue-1; pressDuration++ {
			distanceReached := pressDuration * (timeValue - pressDuration)

			if distanceReached > distanceValue {
				possibleWins++
			}
		}

		victories = append(victories, possibleWins)
	}

	result := victories[0]

	for _, value := range victories[1:] {
		result *= value
	}

	return result
}

func part2(input string) int {
	inputLines := strings.Split(input, "\n")

	time := strings.Join(strings.Fields(strings.Split(inputLines[0], ":")[1]), "")
	distance := strings.Join(strings.Fields(strings.Split(inputLines[1], ":")[1]), "")

	timeValue, _ := strconv.Atoi(time)
	distanceValue, _ := strconv.Atoi(distance)

	possibleWins := 0

	for pressDuration := 1; pressDuration < timeValue-1; pressDuration++ {
		distanceReached := pressDuration * (timeValue - pressDuration)

		if distanceReached > distanceValue {
			possibleWins++
		}
	}

	return possibleWins
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

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

type Mapping struct {
	mapsTo   string
	mappings [][]int
}

func part1(input string) int {
	inputLines := strings.Split(input, "\n")
	if len(inputLines) == 0 {
		panic("malformed input")
	}

	seeds := strings.Fields(strings.Split(inputLines[0], ":")[1])

	almanach := map[string]Mapping{
		"seed": {
			mapsTo: "soil",
		},
		"soil": {
			mapsTo: "fertilizer",
		},
		"fertilizer": {
			mapsTo: "water",
		},
		"water": {
			mapsTo: "light",
		},
		"light": {
			mapsTo: "temperature",
		},
		"temperature": {
			mapsTo: "humidity",
		},
		"humidity": {
			mapsTo: "location",
		},
	}

	parseAlmanach(almanach, inputLines[1:])

	lowestLocation := math.MaxInt

	for _, seed := range seeds {
		currentValue, _ := strconv.Atoi(seed)
		currentMapping := "seed"

		for currentMapping != "location" {
			foundMapping := false

			for _, mapping := range almanach[currentMapping].mappings {
				if currentValue >= mapping[1] && currentValue < mapping[1]+mapping[2] {
					currentValue = mapping[0] + currentValue - mapping[1]
					foundMapping = true
				}

				if foundMapping {
					break
				}
			}

			currentMapping = almanach[currentMapping].mapsTo
		}

		lowestLocation = min(lowestLocation, currentValue)
	}

	return lowestLocation
}

func parseAlmanach(mappings map[string]Mapping, inputLines []string) {
	currentProcessedMapping := ""

	for _, line := range inputLines {
		if len(line) == 0 {
			continue
		}

		currentProcessedMappingHasChanged := false

		for key := range mappings {
			if strings.HasPrefix(line, key) {
				currentProcessedMapping = key
				currentProcessedMappingHasChanged = true
				break
			}
		}

		if currentProcessedMappingHasChanged {
			continue
		}

		mapEntry := mappings[currentProcessedMapping]

		digitsStr := strings.Fields(line)
		var digits []int

		for _, v := range digitsStr {
			converted, _ := strconv.Atoi(v)
			digits = append(digits, converted)
		}

		mapEntry.mappings = append(mapEntry.mappings, digits)
		mappings[currentProcessedMapping] = mapEntry
	}
}

func part2(input string) int {
	// Slow solution, took 8min55s on a M1 Pro.
	// Will work on a scalable algorithm later.
	inputLines := strings.Split(input, "\n")
	if len(inputLines) == 0 {
		panic("malformed input")
	}

	seeds := strings.Fields(strings.Split(inputLines[0], ":")[1])

	almanach := map[string]Mapping{
		"seed": {
			mapsTo: "soil",
		},
		"soil": {
			mapsTo: "fertilizer",
		},
		"fertilizer": {
			mapsTo: "water",
		},
		"water": {
			mapsTo: "light",
		},
		"light": {
			mapsTo: "temperature",
		},
		"temperature": {
			mapsTo: "humidity",
		},
		"humidity": {
			mapsTo: "location",
		},
	}

	parseAlmanach(almanach, inputLines[1:])

	lowestLocation := math.MaxInt

	var newSeedsRanges [][]string
	for i := 0; i < len(seeds); i = i + 2 {
		newSeedsRanges = append(newSeedsRanges, []string{seeds[i], seeds[i+1]})
	}

	for _, seedRange := range newSeedsRanges {
		seedRangeStart, _ := strconv.Atoi(seedRange[0])
		seedRangeLength, _ := strconv.Atoi(seedRange[1])

		for currentSeed := seedRangeStart; currentSeed < seedRangeStart+seedRangeLength; currentSeed++ {
			currentValue := currentSeed
			currentMapping := "seed"

			for currentMapping != "location" {
				foundMapping := false

				for _, mapping := range almanach[currentMapping].mappings {
					if currentValue >= mapping[1] && currentValue < mapping[1]+mapping[2] {
						currentValue = mapping[0] + currentValue - mapping[1]
						foundMapping = true
					}

					if foundMapping {
						break
					}
				}

				currentMapping = almanach[currentMapping].mapsTo
			}

			lowestLocation = min(lowestLocation, currentValue)
		}

	}

	return lowestLocation
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

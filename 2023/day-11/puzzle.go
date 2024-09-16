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
	sum := 0

	inputLines := strings.Split(input, "\n")

	emptyLines := map[int]bool{}
	emptyColumns := map[int]bool{}
	notEmptyColumns := map[int]bool{}

	galaxyMap := map[int][]int{}

	originalWidth := len(inputLines[0])
	originalHeight := 0

	galaxyNameGenerator := 0

	for indexLine, line := range inputLines {
		if line == "" {
			continue
		}

		originalHeight++

		foundGalaxyInLine := false
		for indexColumn, scanned := range line {
			if scanned == '#' {
				foundGalaxyInLine = true
				notEmptyColumns[indexColumn] = true
				galaxyMap[galaxyNameGenerator] = []int{indexLine, indexColumn}
				galaxyNameGenerator++
			}
		}
		if !foundGalaxyInLine {
			emptyLines[indexLine] = true
		}
	}

	for i := 0; i < originalWidth; i++ {
		if !notEmptyColumns[i] {
			emptyColumns[i] = true
		}
	}

	scannedPairs := map[int]map[int]bool{}
	galaxiesFounds := len(galaxyMap)

	for nameGalaxy1, galaxy1 := range galaxyMap {
		for nameGalaxy2, galaxy2 := range galaxyMap {
			if nameGalaxy1 == nameGalaxy2 {
				continue
			}

			smallerName := min(nameGalaxy1, nameGalaxy2)
			biggerName := max(nameGalaxy1, nameGalaxy2)

			if len(scannedPairs[smallerName]) >= galaxiesFounds-1 || scannedPairs[smallerName][biggerName] == true {
				continue
			}

			smallerX := min(galaxy1[1], galaxy2[1])
			biggerX := max(galaxy1[1], galaxy2[1])
			distanceX := biggerX - smallerX

			if distanceX > 1 {
				for i := smallerX + 1; i < biggerX; i++ {
					if emptyColumns[i] {
						distanceX++
					}
				}
			}

			smallerY := min(galaxy1[0], galaxy2[0])
			biggerY := max(galaxy1[0], galaxy2[0])
			distanceY := biggerY - smallerY

			if distanceY > 1 {
				for i := smallerY + 1; i < biggerY; i++ {
					if emptyLines[i] {
						distanceY++
					}
				}
			}

			if scannedPairs[smallerName] == nil {
				scannedPairs[smallerName] = map[int]bool{}
			}
			scannedPairs[smallerName][biggerName] = true

			sum += distanceX + distanceY
		}
	}

	return sum
}

func part2(input string) int {
	sum := 0

	expansionFactor := 1000000

	inputLines := strings.Split(input, "\n")

	emptyLines := map[int]bool{}
	emptyColumns := map[int]bool{}
	notEmptyColumns := map[int]bool{}

	galaxyMap := map[int][]int{}

	originalWidth := len(inputLines[0])
	originalHeight := 0

	galaxyNameGenerator := 0

	for indexLine, line := range inputLines {
		if line == "" {
			continue
		}

		originalHeight++

		foundGalaxyInLine := false
		for indexColumn, scanned := range line {
			if scanned == '#' {
				foundGalaxyInLine = true
				notEmptyColumns[indexColumn] = true
				galaxyMap[galaxyNameGenerator] = []int{indexLine, indexColumn}
				galaxyNameGenerator++
			}
		}
		if !foundGalaxyInLine {
			emptyLines[indexLine] = true
		}
	}

	for i := 0; i < originalWidth; i++ {
		if !notEmptyColumns[i] {
			emptyColumns[i] = true
		}
	}

	scannedPairs := map[int]map[int]bool{}
	galaxiesFounds := len(galaxyMap)

	for nameGalaxy1, galaxy1 := range galaxyMap {
		for nameGalaxy2, galaxy2 := range galaxyMap {
			if nameGalaxy1 == nameGalaxy2 {
				continue
			}

			smallerName := min(nameGalaxy1, nameGalaxy2)
			biggerName := max(nameGalaxy1, nameGalaxy2)

			if len(scannedPairs[smallerName]) >= galaxiesFounds-1 || scannedPairs[smallerName][biggerName] == true {
				continue
			}

			smallerX := min(galaxy1[1], galaxy2[1])
			biggerX := max(galaxy1[1], galaxy2[1])
			distanceX := biggerX - smallerX

			if distanceX > 1 {
				for i := smallerX + 1; i < biggerX; i++ {
					if emptyColumns[i] {
						distanceX += expansionFactor - 1
					}
				}
			}

			smallerY := min(galaxy1[0], galaxy2[0])
			biggerY := max(galaxy1[0], galaxy2[0])
			distanceY := biggerY - smallerY

			if distanceY > 1 {
				for i := smallerY + 1; i < biggerY; i++ {
					if emptyLines[i] {
						distanceY += expansionFactor - 1
					}
				}
			}

			if scannedPairs[smallerName] == nil {
				scannedPairs[smallerName] = map[int]bool{}
			}
			scannedPairs[smallerName][biggerName] = true

			sum += distanceX + distanceY
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

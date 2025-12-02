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

type NodeNext struct {
	isStart bool
	isEnd   bool
	left    string
	right   string
}

func part1(input string) int {
	steps := 0

	parsedMap := map[string]NodeNext{}

	inputLines := strings.Split(input, "\n")

	pattern := []rune{}
	for _, c := range inputLines[0] {
		pattern = append(pattern, c)
	}
	patternLength := len(pattern)

	for _, line := range inputLines[1:] {
		if line == "" {
			continue
		}

		explodedLine := strings.Split(line, "=")
		current := strings.TrimSpace(explodedLine[0])
		choice := strings.Split(explodedLine[1], ",")
		left := strings.Trim(choice[0], " (")
		right := strings.Trim(choice[1], " )")

		parsedMap[current] = NodeNext{
			left:  left,
			right: right,
		}
	}

	currentNode := "AAA"
	i := 0
	for i < patternLength {
		var next string
		if pattern[i] == 'L' {
			next = parsedMap[currentNode].left
		} else {
			next = parsedMap[currentNode].right
		}

		steps++

		if next == "ZZZ" {
			break
		}

		currentNode = next
		i = (i + 1) % patternLength
	}

	return steps
}

type EndlessEight struct {
	startingNode         string
	currentProcessedNode string
	distanceToEnding     int
	endingNode           string
	loopLength           int
	evaluatingLoop       bool
}

func part2(input string) int {
	steps := 0

	parsedMap := map[string]NodeNext{}

	startingNodes := map[string]EndlessEight{}

	inputLines := strings.Split(input, "\n")

	pattern := []rune{}
	for _, c := range inputLines[0] {
		pattern = append(pattern, c)
	}
	patternLength := len(pattern)

	for _, line := range inputLines[1:] {
		if line == "" {
			continue
		}

		explodedLine := strings.Split(line, "=")
		current := strings.TrimSpace(explodedLine[0])

		choice := strings.Split(explodedLine[1], ",")
		left := strings.Trim(choice[0], " (")
		right := strings.Trim(choice[1], " )")

		isEnd := false

		switch current[2] {
		case 'A':
			startingNodes[current] = EndlessEight{
				startingNode:         current,
				currentProcessedNode: current,
			}
		case 'Z':
			isEnd = true
		}

		parsedMap[current] = NodeNext{
			isEnd: isEnd,
			left:  left,
			right: right,
		}
	}

	startingNodesLength := len(startingNodes)
	i := 0
	for i < patternLength {
		steps++

		for _, currentNode := range startingNodes {
			var next string
			if pattern[i] == 'L' {
				next = parsedMap[currentNode.currentProcessedNode].left
			} else {
				next = parsedMap[currentNode.currentProcessedNode].right
			}

			if parsedMap[next].isEnd {
				if currentNode.distanceToEnding == 0 {
					currentNode.distanceToEnding = steps
				}
				if currentNode.endingNode == "" {
					currentNode.endingNode = next
					currentNode.evaluatingLoop = true
				} else {
					if currentNode.evaluatingLoop {
						currentNode.loopLength++
						currentNode.evaluatingLoop = false
					}
				}
			}

			currentNode.currentProcessedNode = next
		}

		hasFinishedAnalysis := 0
		for _, currentNode := range startingNodes {
			if currentNode.loopLength > 0 && currentNode.evaluatingLoop == false {
				hasFinishedAnalysis++
			}
		}

		if hasFinishedAnalysis == startingNodesLength {
			continue
		}

		i = (i + 1) % patternLength
	}

	loopCount := map[string]int{}

	// Magic formula to calculate loops

	results := 0
	for key, currentNode := range startingNodes {
		results += currentNode.distanceToEnding + loopCount[key]*currentNode.loopLength
	}

	return steps
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

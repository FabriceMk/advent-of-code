package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"sort"
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

type HandBid struct {
	hand                    []int
	handValueRepresentation int
	bid                     int
}

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

// Today's implementation uis super ugly with a dirty black hack to sort same hands types by comparing
// an integer representation that keeps the ordering of a hand.
// The first implementation tried to sort arrays at each insert of a hand but it was super slow for some reason.
// Need to delete the hack while keeping the speed.
func part1(input string) int {
	sum := 0
	rank := 0

	inputLines := strings.Split(input, "\n")

	// Will have a size of 7, one per type of hand. If multiple hands have the same type,
	// they are sorted from the strongest to the weakest.
	rankingHands := make([][]HandBid, 7)

	for _, line := range inputLines {
		rank++

		cardsCount := map[int]int{}

		lineSplit := strings.Split(line, " ")

		hand := lineSplit[0]
		bid, _ := strconv.Atoi(lineSplit[1])

		handBid := HandBid{
			bid: bid,
		}

		// Multiplicator to get an integer representation of the value of a hand for sorting
		// Dirty hack to remove later
		multiplicator := 100 * 100 * 100 * 100 * 100
		for _, card := range hand {
			cardValue := getCardPower(card)
			handBid.hand = append(handBid.hand, cardValue)

			// Dirty hack
			handBid.handValueRepresentation += cardValue * multiplicator
			multiplicator /= 100
			if multiplicator == 0 {
				multiplicator = 1
			}

			cardsCount[cardValue] += 1
		}

		handPower := 0

		counts := []int{}
		for _, count := range cardsCount {
			counts = append(counts, count)
		}
		slices.Sort(counts)

		countStrRepresentation := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(counts)), ""), "[]")

		switch countStrRepresentation {
		case "5":
			handPower = FIVE_OF_A_KIND
		case "14":
			handPower = FOUR_OF_A_KIND
		case "23":
			handPower = FULL_HOUSE
		case "113":
			handPower = THREE_OF_A_KIND
		case "122":
			handPower = TWO_PAIR
		case "1112":
			handPower = ONE_PAIR
		default:
			handPower = HIGH_CARD
		}

		rankingHands[handPower] = append(rankingHands[handPower], handBid)
	}

	for i := FIVE_OF_A_KIND; i >= HIGH_CARD; i-- {
		if len(rankingHands[i]) == 0 {
			continue
		}

		sameTypeHands := rankingHands[i]
		sort.Slice(sameTypeHands, func(i, j int) bool {
			return sameTypeHands[i].handValueRepresentation > sameTypeHands[j].handValueRepresentation
		})

		for _, handbid := range rankingHands[i] {
			sum += rank * handbid.bid
			rank--
		}
	}

	return sum
}

func getCardPower(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		value, _ := strconv.Atoi(string(card))
		return value
	}
}

func isHandStronger(hand1 []int, hand2 []int) bool {
	for i := 0; i < len(hand1); i++ {
		if hand1[i] == hand2[i] {
			continue
		}
		return hand1[i] > hand2[i]
	}

	return false
}

func part2(input string) int {
	sum := 0
	rank := 0

	const JOKER_VALUE = 1

	inputLines := strings.Split(input, "\n")

	// Will have a size of 7, one per type of hand. If multiple hands have the same type,
	// they are sorted from the strongest to the weakest.
	rankingHands := make([][]HandBid, 7)

	for _, line := range inputLines {
		rank++

		cardsCount := make(map[int]int)

		lineSplit := strings.Split(line, " ")

		hand := lineSplit[0]
		bid, _ := strconv.Atoi(lineSplit[1])

		handBid := HandBid{
			bid: bid,
		}

		// Multiplicator to get an integer representation of the value of a hand for sorting
		// Dirty hack to remove later
		multiplicator := 100 * 100 * 100 * 100 * 100
		jokersCount := 0
		for _, card := range hand {
			if card == 'J' {
				jokersCount++
			}
			cardValue := getCardPowerAlternate(card)
			handBid.hand = append(handBid.hand, cardValue)

			// Dirty hack
			handBid.handValueRepresentation += cardValue * multiplicator
			multiplicator /= 100
			if multiplicator == 0 {
				multiplicator = 1
			}

			cardsCount[cardValue] += 1
		}

		if jokersCount > 0 {
			highestCount := 0
			cardValueHighestCount := 0

			for cardValue, count := range cardsCount {
				if cardValue != JOKER_VALUE && count > highestCount {
					highestCount = count
					cardValueHighestCount = cardValue
				}
			}

			cardsCount[cardValueHighestCount] += jokersCount
			delete(cardsCount, JOKER_VALUE)
		}

		counts := []int{}
		for _, count := range cardsCount {
			counts = append(counts, count)
		}
		slices.Sort(counts)

		handPower := 0
		countStrRepresentation := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(counts)), ""), "[]")

		switch countStrRepresentation {
		case "5":
			handPower = FIVE_OF_A_KIND
		case "14":
			handPower = FOUR_OF_A_KIND
		case "23":
			handPower = FULL_HOUSE
		case "113":
			handPower = THREE_OF_A_KIND
		case "122":
			handPower = TWO_PAIR
		case "1112":
			handPower = ONE_PAIR
		default:
			handPower = HIGH_CARD
		}

		rankingHands[handPower] = append(rankingHands[handPower], handBid)
	}

	for i := FIVE_OF_A_KIND; i >= HIGH_CARD; i-- {
		if len(rankingHands[i]) == 0 {
			continue
		}

		sameTypeHands := rankingHands[i]
		sort.Slice(sameTypeHands, func(i, j int) bool {
			return sameTypeHands[i].handValueRepresentation > sameTypeHands[j].handValueRepresentation
		})

		for _, handbid := range rankingHands[i] {
			sum += rank * handbid.bid
			rank--
		}
	}

	return sum
}

func getCardPowerAlternate(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		value, _ := strconv.Atoi(string(card))
		return value
	}
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

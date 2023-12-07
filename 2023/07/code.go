package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type hand struct {
	cards string
	bid   int
	score int64
}

var (
	hands           []hand
	part1CardValues = map[rune]int64{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	part2CardValues = map[rune]int64{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
)

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	processInput(input, part2)

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].score > hands[j].score
	})
	sum := 0
	handLen := len(hands)
	for i, hand := range hands {
		sum += hand.bid * (handLen - i)
	}
	return sum
}

func scoreHand(cards string, part2 bool) int64 {
	score := int64(0)
	counts := make(map[rune]int)
	for i, card := range cards {
		counts[card]++
		if part2 {
			score += part2CardValues[card] * int64(math.Pow(10, float64((4-i)*2)))
		} else {
			score += part1CardValues[card] * int64(math.Pow(10, float64((4-i)*2)))
		}
	}
	highest := 0
	pairCount := 0
	for card, count := range counts {
		if part2 && card == 'J' {
			continue
		}
		if count > highest {
			highest = count
		}
		if count == 2 {
			pairCount++
		}
	}
	if part2 {
		if highest == 2 {
			pairCount--
		}
		highest += counts['J']
	}
	switch highest {
	case 5:
		score += 70000000000
	case 4:
		score += 60000000000
	case 3:
		if pairCount == 1 {
			score += 50000000000
		} else {
			score += 40000000000
		}
	case 2:
		if pairCount == 2 {
			score += 30000000000
		} else {
			score += 20000000000
		}
	case 1:
		score += 10000000000
	}
	return score
}

func processInput(input string, part2 bool) {
	hands = make([]hand, 0)
	for _, line := range strings.Split(input, "\n") {
		before, after, found := strings.Cut(line, " ")
		if !found {
			panic("invalid input")
		}
		bid, err := strconv.Atoi(after)
		if err != nil {
			panic(err)
		}
		hands = append(hands, hand{
			cards: before,
			bid:   bid,
			score: scoreHand(before, part2),
		})
	}
}

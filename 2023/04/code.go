package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type card struct {
	winningNumbers map[int]bool
	heldNumbers    []int
}

func newCard(winning map[int]bool, held []int) *card {
	return &card{
		winningNumbers: winning,
		heldNumbers:    held,
	}
}

var cards map[int]*card

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	processInput(input)
	if part2 {
		return part2fn()
	}
	// solve part 1 here
	return part1fn()
}

func part1fn() int {
	sum := 0
	for _, c := range cards {
		n := 0
		for _, h := range c.heldNumbers {
			if c.winningNumbers[h] {
				n++
			}
		}
		if n > 0 {
			// 2 ^ (n-1)
			sum += int(math.Pow(2, float64(n-1)))
		}
	}
	return sum
}

func part2fn() int {
	cardsLen := len(cards)
	sum := 0
	copies := make([]int, cardsLen)

	var c *card
	for i := 1; i <= cardsLen; i++ {
		sum += copies[i-1] + 1
		c = cards[i]
		n := 0
		for _, h := range c.heldNumbers {
			if c.winningNumbers[h] {
				n++
			}
		}
		if n > 0 {
			for j := 0; j < n; j++ {
				if i+j >= cardsLen {
					break
				}
				copies[i+j] += copies[i-1] + 1
			}
		}
	}

	return sum
}

func processInput(input string) {
	cards = make(map[int]*card)

	startI := 0

	capturing := false
	wantWinning := false
	wantHeld := false

	id := 1
	winningNumbers := make(map[int]bool)
	heldNumbers := make([]int, 0, 25)

	for i, r := range input {
		if r < '0' || r > '9' {
			if capturing {
				num, err := strconv.Atoi(input[startI:i])
				if err != nil {
					panic(fmt.Sprintf("invalid input of number: %s", input[startI:i]))
				}
				if wantWinning {
					winningNumbers[num] = true
				} else if wantHeld {
					heldNumbers = append(heldNumbers, num)
				}
				capturing = false
			}
			if r == '\n' {
				wantWinning = false
				wantHeld = false
				cards[id] = newCard(winningNumbers, heldNumbers)
				id++
				winningNumbers = make(map[int]bool)
				heldNumbers = make([]int, 0, 25)
			}
		} else {
			if !capturing {
				startI = i
				capturing = true
			}
			continue
		}

		if r == ':' {
			wantWinning = true
			continue
		} else if r == '|' {
			wantWinning = false
			wantHeld = true
			continue
		}
	}

	if capturing {
		num, err := strconv.Atoi(input[startI:])
		if err != nil {
			panic(fmt.Sprintf("invalid input of number: %s", input[startI:]))
		}
		if wantWinning {
			winningNumbers[num] = true
		} else if wantHeld {
			heldNumbers = append(heldNumbers, num)
		}
		capturing = false
	}
	wantWinning = false
	wantHeld = false
	cards[id] = newCard(winningNumbers, heldNumbers)
	id++
	winningNumbers = make(map[int]bool)
	heldNumbers = make([]int, 0, 25)
}

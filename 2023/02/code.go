package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if part2 {
			sum += processLine2(line)
		} else {
			sum += processLine(line)
		}
	}
	return sum
}

func processLine(input string) int {
	valid := make(map[string]int)
	valid["red"] = 12
	valid["green"] = 13
	valid["blue"] = 14

	b, a, f := strings.Cut(input[5:], ":")
	if !f {
		fmt.Println("bad input:", input)
		return 0
	}

	id, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("id not number:", input)
		return 0
	}

	for _, round := range strings.Split(a[1:], "; ") {
		for _, check := range strings.Split(round, ", ") {
			cb, ca, cf := strings.Cut(check, " ")
			if !cf {
				fmt.Println("check doesn't have space:", check)
				return 0
			}

			c, err := strconv.Atoi(cb)
			if err != nil {
				fmt.Println("check count not number:", check)
				return 0
			}

			if valid[ca] < c {
				return 0
			}
		}
	}

	return id
}

func processLine2(input string) int {
	valid := make(map[string]int)

	b, a, f := strings.Cut(input[5:], ":")
	if !f {
		fmt.Println("bad input:", input)
		return 0
	}

	_, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("id not number:", input)
		return 0
	}

	for _, round := range strings.Split(a[1:], "; ") {
		for _, check := range strings.Split(round, ", ") {
			cb, ca, cf := strings.Cut(check, " ")
			if !cf {
				fmt.Println("check doesn't have space:", check)
				return 0
			}

			c, err := strconv.Atoi(cb)
			if err != nil {
				fmt.Println("check count not number:", check)
				return 0
			}

			if valid[ca] < c {
				valid[ca] = c
			}
		}
	}

	return valid["red"] * valid["green"] * valid["blue"]
}

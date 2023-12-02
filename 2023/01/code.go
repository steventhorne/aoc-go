package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var notANumberError = errors.New("not a number")

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if part2 {
			sum += processLine(part2, line)
		} else {
			sum += processLine(part2, line)
		}
	}
	return sum
}

func convertTextNumber(input string) (string, error) {
	if len(input) >= 3 && input[:3] == "one" {
		return "1", nil
	}
	if len(input) >= 3 && input[:3] == "two" {
		return "2", nil
	}
	if len(input) >= 5 && input[:5] == "three" {
		return "3", nil
	}
	if len(input) >= 4 && input[:4] == "four" {
		return "4", nil
	}
	if len(input) >= 4 && input[:4] == "five" {
		return "5", nil
	}
	if len(input) >= 3 && input[:3] == "six" {
		return "6", nil
	}
	if len(input) >= 5 && input[:5] == "seven" {
		return "7", nil
	}
	if len(input) >= 5 && input[:5] == "eight" {
		return "8", nil
	}
	if len(input) >= 4 && input[:4] == "nine" {
		return "9", nil
	}
	if len(input) >= 4 && input[:4] == "zero" {
		return "0", nil
	}
	return "", notANumberError
}

func processLine(part2 bool, line string) int {
	first := ""
	cur := ""

	for i, c := range line {
		if c >= '0' && c <= '9' {
			cur = string(c)
		} else if part2 {
			if tn, err := convertTextNumber(line[i:]); err == nil {
				cur = tn
			}
		}

		if first == "" && cur != "" {
			first = cur
		}
	}

	if val, err := strconv.Atoi(first + cur); err == nil {
		return val
	}
	return 0
}

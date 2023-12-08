package main

import (
	"aoc-in-go/utils"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var (
	sliceLens = map[string]int{
		"L": 3,
		"R": 4,
	}
)

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return part2fn(input)
	}
	// solve part 1 here
	return part1fn(input)
}

func part1fn(input string) (steps int) {
	steps = 0
	currentStep := "AAA"
	dirs, maps := processInput(input)

	for {
		for _, dir := range dirs {
			currentStep = maps[dir][currentStep]
			steps++
			if currentStep == "ZZZ" {
				return steps
			}
		}
	}
}

// find least common multiple of all steps from each start to end
func part2fn(input string) (totalSteps int64) {
	dirs, maps := processInput(input)
	starts := make([]string, 0)

	for step := range maps['L'] {
		if step[2] == 'A' {
			starts = append(starts, step)
		}
	}

	steps := make([]int64, len(starts))

	for i, start := range starts {
	startLoop:
		for {
			for _, dir := range dirs {
				start = maps[dir][start]
				steps[i]++
				if start[2] == 'Z' {
					break startLoop
				}
			}
		}
	}

	return utils.Lcm(steps...)
}

func processInput(input string) (dirs string, maps map[rune]map[string]string) {
	lines := strings.Split(input, "\n")
	dirs = lines[0]

	maps = make(map[rune]map[string]string)
	maps['L'] = make(map[string]string)
	maps['R'] = make(map[string]string)
	for _, line := range lines[2:] {
		maps['L'][line[:3]] = line[7:10]
		maps['R'][line[:3]] = line[12:15]
	}
	return dirs, maps
}

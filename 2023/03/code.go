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

type direction struct {
	x int
	y int
}

var (
	grid       [140][140]rune
	used       [140][140]bool
	directions = []direction{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
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
	parseInput(input)
	used = [140][140]bool{}
	if part2 {
		return sumGearRatios()
	}
	return sumParts()
}

func parseInput(input string) {
	grid = [140][140]rune{}
	for y, line := range strings.Split(input, "\n") {
		for x, r := range line {
			grid[y][x] = r
		}
	}
}

func sumParts() int {
	sum := 0
	for y, line := range grid {
		for x, r := range line {
			// check if symbol
			if (r < '0' || r > '9') && r != '.' && r != 0 {
				for _, d := range directions {
					val, err := getInnerNumber(x+d.x, y+d.y)
					if err == nil {
						sum += val
					}
				}
			}
		}
	}
	return sum
}

func sumGearRatios() int {
	sum := 0
	for y, line := range grid {
		for x, r := range line {
			// check if gear symbol
			if r == '*' {
				gearRatios := make([]int, 0)
				var mult int
				for _, d := range directions {
					val, err := getInnerNumber(x+d.x, y+d.y)
					if err == nil {
						gearRatios = append(gearRatios, val)
					}
				}
				if len(gearRatios) > 1 {
					mult = 1
					for _, v := range gearRatios {
						mult *= v
					}
					sum += mult
				}
			}
		}
	}
	return sum
}

var (
	outOfRange  = errors.New("index out of range")
	notANumber  = errors.New("not a number")
	alreadyUsed = errors.New("already used")
)

func getInnerNumber(x int, y int) (int, error) {
	if x < 0 || x > 139 || y < 0 || y > 139 {
		return 0, outOfRange
	}
	gridLine := grid[y]
	if gridLine[x] < '0' || gridLine[x] > '9' {
		return 0, notANumber
	}

	earliest := x
	latest := x
	for earliest > 0 && gridLine[earliest-1] >= '0' && gridLine[earliest-1] <= '9' {
		earliest--
	}
	for latest < 139 && gridLine[latest+1] >= '0' && gridLine[latest+1] <= '9' {
		latest++
	}
	if used[y][earliest] {
		return 0, alreadyUsed
	}
	used[y][earliest] = true
	num, err := strconv.Atoi(string(gridLine[earliest : latest+1]))
	if err != nil {
		panic(err)
	}
	return num, nil
}

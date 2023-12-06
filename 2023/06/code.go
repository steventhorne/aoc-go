package main

import (
	"math"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var (
	times         []int
	distances     []int
	numberStrings = map[rune]string{
		'0': "0",
		'1': "1",
		'2': "2",
		'3': "3",
		'4': "4",
		'5': "5",
		'6': "6",
		'7': "7",
		'8': "8",
		'9': "9",
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
		processInput2(input)
		return solution()
	}
	// solve part 1 here
	processInput1(input)
	return solution()
}

func solution() int {
	mult := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		chargeTime := getChargeTimeToBeat(time, distance)
		waysToWin := int(((float64(time)/2)-float64(chargeTime))*2 + 1)
		mult *= waysToWin
	}
	return mult
}

func processInput1(input string) {
	time := true
	times = make([]int, 0)
	distances = make([]int, 0)

	numStr := ""
	for i, r := range input {
		numStr += numberStrings[r]
		if r == '\n' || r == ' ' || i == len(input)-1 {
			if numStr != "" {
				val, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				if time {
					times = append(times, val)
				} else {
					distances = append(distances, val)
				}
				numStr = ""
			}
			if r == '\n' {
				time = false
			}
		}
	}
}

func processInput2(input string) {
	time := true
	times = make([]int, 0)
	distances = make([]int, 0)

	numStr := ""
	for i, r := range input {
		numStr += numberStrings[r]
		if r == '\n' || i == len(input)-1 {
			if numStr != "" {
				val, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				if time {
					times = append(times, val)
				} else {
					distances = append(distances, val)
				}
				numStr = ""
			}
			if r == '\n' {
				time = false
			}
		}
	}
}

func getChargeTimeToBeat(totalTime int, timeToBeat int) int {
	var a = -(-float64(totalTime) + math.Sqrt(float64((totalTime*totalTime)-4*timeToBeat))) / 2.0
	final := int(math.Ceil(a))
	if (totalTime-final)*final == timeToBeat {
		return final + 1
	}
	return final
}

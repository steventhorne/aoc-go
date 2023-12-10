package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var (
	numsList [][]int
)

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	processInput(input)
	return solve(part2)
}

func solve(part2 bool) int {
	firsts := make([]int, len(numsList[0]))
	sum := 0
	for _, nums := range numsList {
		nl := len(nums)
		for nl > 0 {
			numSlice := nums[:nl]
			nl--
			firsts[nl] = numSlice[0]

			allZeros := true
			r := numSlice[nl]
			l := 0
			for i := nl; i > 0; i-- {
				l = numSlice[i-1]
				numSlice[i-1] = r - l
				if numSlice[i-1] != 0 {
					allZeros = false
				}
				r = l
			}

			if allZeros {
				if part2 {
					miniSum := firsts[nl]
					for i := nl; i < len(firsts)-1; i++ {
						miniSum = firsts[i+1] - miniSum
					}
					sum += miniSum
				} else {
					for i := nl; i < len(nums); i++ {
						sum += nums[i]
					}
				}
				break
			}
		}
	}
	return sum
}

func processInput(input string) {
	numsList = [][]int{}
	for _, line := range strings.Split(input, "\n") {
		nums := []int{}
		for _, num := range strings.Split(line, " ") {
			num, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		numsList = append(numsList, nums)
	}
}

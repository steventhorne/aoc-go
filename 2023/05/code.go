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
	seeds    []int64
	farmMaps = make(map[int]*farmMap)
)

const (
	sectionSeeds = iota
	sectionToSoil
	sectionToFertilizer
	sectionToWater
	sectionToLight
	sectionToTemperature
	sectionToHumidity
	sectionToLocation
	sectionSize
)

type farmMap struct {
	ranges []mapping
}

func newFarmMap() *farmMap {
	return &farmMap{
		ranges: make([]mapping, 0),
	}
}

type mapping struct {
	destStart   int64
	sourceStart int64
	size        int64
}

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

func part1fn() int64 {
	location := int64(-1)
	for _, seed := range seeds {
		curVal := seed
		for i := sectionSeeds; i < sectionSize; i++ {
			farmMap := farmMaps[i]
			for _, mapping := range farmMap.ranges {
				if curVal >= mapping.sourceStart && curVal < mapping.sourceStart+mapping.size {
					curVal = mapping.destStart + (curVal - mapping.sourceStart)
					break
				}
			}
		}
		if location < 0 || curVal < location {
			location = curVal
		}
	}

	return location
}

func part2fn() int64 {
	location := int64(-1)
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed <= seeds[i]+seeds[i+1]; seed++ {
			curVal := seed
			for i := sectionSeeds; i < sectionSize; i++ {
				farmMap := farmMaps[i]
				for _, mapping := range farmMap.ranges {
					if curVal >= mapping.sourceStart && curVal < mapping.sourceStart+mapping.size {
						curVal = mapping.destStart + (curVal - mapping.sourceStart)
						break
					}
				}
			}
			if location < 0 || curVal < location {
				location = curVal
			}
		}
	}

	return location
}

func processInput(input string) {
	section := sectionSeeds
	farmMaps[section] = newFarmMap()

	skipLine := 0
	for _, line := range strings.Split(input, "\n") {
		if skipLine > 0 {
			skipLine--
			continue
		}
		switch section {
		case sectionSeeds:
			seedsStr := strings.Split(line, " ")
			seeds = make([]int64, len(seedsStr[1:]))
			for i, seedStr := range seedsStr[1:] {
				val, err := strconv.ParseInt(seedStr, 10, 64)
				if err != nil {
					panic(err)
				}

				seeds[i] = val
			}
			section = sectionToSoil
			farmMaps[section] = newFarmMap()
			skipLine = 2
		case sectionToSoil:
			if line == "" {
				section = sectionToFertilizer
				farmMaps[section] = newFarmMap()
				skipLine = 1
				continue
			}
			farmMaps[section].ranges = append(farmMaps[section].ranges, getMapping(line))
		case sectionToFertilizer:
			if line == "" {
				section = sectionToWater
				farmMaps[section] = newFarmMap()
				skipLine = 1
				continue
			}
			farmMaps[section].ranges = append(farmMaps[section].ranges, getMapping(line))
		case sectionToWater:
			if line == "" {
				section = sectionToLight
				farmMaps[section] = newFarmMap()
				skipLine = 1
				continue
			}
			farmMaps[section].ranges = append(farmMaps[section].ranges, getMapping(line))
		case sectionToLight:
			if line == "" {
				section = sectionToTemperature
				farmMaps[section] = newFarmMap()
				skipLine = 1
				continue
			}
			farmMaps[section].ranges = append(farmMaps[section].ranges, getMapping(line))
		case sectionToTemperature:
			if line == "" {
				section = sectionToHumidity
				farmMaps[section] = newFarmMap()
				skipLine = 1
				continue
			}
			farmMaps[section].ranges = append(farmMaps[section].ranges, getMapping(line))
		case sectionToHumidity:
			if line == "" {
				section = sectionToLocation
				farmMaps[section] = newFarmMap()
				skipLine = 1
				continue
			}
			farmMaps[section].ranges = append(farmMaps[section].ranges, getMapping(line))
		case sectionToLocation:
			if line == "" {
				return
			}
			farmMaps[section].ranges = append(farmMaps[section].ranges, getMapping(line))
		}
	}
}

func getMapping(line string) mapping {
	parts := strings.Split(line, " ")
	destStart, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(err)
	}
	sourceStart, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}
	size, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		panic(err)
	}
	return mapping{
		destStart:   destStart,
		sourceStart: sourceStart,
		size:        size,
	}
}

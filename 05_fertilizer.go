package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Day 04 solution
func fertilizer(filename string, calcFun func([]string) int, prepFun func(data []string) []string) int {
	input := readInput(filename, calcFun, prepFun)

	return calcFun(input)
}

// Solve puzzle no 1
func puzzle1(input []string) int {
	var res int = -1

	translated := false
	seeds := splitToInts(input[0])

	for j := 0; j < len(seeds); j++ {
		fmt.Println("Seed: ", seeds[j])
		translated = false

		for i := 2; i < len(input); i++ {
			// fmt.Println("Input: ", input[i])
			if strings.Contains(input[i], "map") {
				fmt.Println("--- Start of map")
				continue
			}
			if input[i] == "" {
				// end of map
				if !translated {
					fmt.Println("No translation found: ", seeds[j], " => saving ", seeds[j])
				}
				translated = false
				fmt.Println("--- End of map\n")
				continue
			}
			// translate
			srcStart, srcEnd, offset := parseRange(input[i])
			// fmt.Println("Src from: ", srcStart, " to: ", srcEnd, " offset ", offset)
			if !translated && srcStart <= seeds[j] && seeds[j] <= srcEnd {
				fmt.Println("Translation found: ", seeds[j], " => ", seeds[j]+offset)
				seeds[j] += offset
				translated = true
			}
		}
		if res == -1 {
			res = seeds[j]
		}
		res = min(res, seeds[j])

	}

	fmt.Println("Result: ", seeds)

	return res
}

func parseRange(line string) (int, int, int) {
	split := splitToInts(line)

	// fmt.Println("Line: ", line, " split: ", split)

	if len(split) != 3 {
		return -1, -1, 0
	}

	// fmt.Println("Src from: ", split[1], " to: ", split[1]+split[2]-1, " offset ", split[0]-split[1])
	return split[1], split[1] + split[2] - 1, split[0] - split[1]
}

// Split text into int values
func splitToInts(text string) []int {
	tmp := strings.Fields(text)
	// fmt.Println(tmp)
	values := make([]int, 0, len(tmp))

	for _, raw := range tmp {
		// fmt.Println(raw)
		v, err := strconv.Atoi(raw)
		if err != nil {
			// log.Print(err)
			continue
		}
		values = append(values, v)
	}
	// fmt.Println("Line: ", text, " split: ", values)

	return values
}

type destination struct {
	srcStart int
	srcEnd   int
	offset   int
}

// Solve puzzle no 2
func puzzle2(input []string) int {
	var res int = -1

	seedsTmp := splitToInts(input[0])

	seeds := make([][]int, 2)
	seeds[0] = []int{seedsTmp[0], seedsTmp[0] + seedsTmp[1] - 1}
	seeds[1] = []int{seedsTmp[2], seedsTmp[2] + seedsTmp[3] - 1}

	// seeds[0] = {}
	// mapResults := map[int]int{}

	// for i := seedsTmp[0]; i < seedsTmp[0]+seedsTmp[1]; i++ {
	// 	seeds = append(seeds, i)
	// 	// mapResults[i] = 0
	// }

	// for i := seedsTmp[2]; i < seedsTmp[2]+seedsTmp[3]; i++ {
	// 	seeds = append(seeds, i)
	// 	// mapResults[i] = 0
	// }

	fmt.Println("Seeds list: ", seeds)

	// read all translations into structure
	translMap := make([][]destination, 7)
	var mapIndex int = 0

	for i := 2; i < len(input); i++ {
		// fmt.Println("Input: ", input[i])
		if strings.Contains(input[i], "map") {
			// fmt.Println("--- Start of map")
			continue
		}
		if input[i] == "" {
			// end of map
			// sort ranges
			sort.Slice(translMap[mapIndex], func(i, j int) bool {
				return translMap[mapIndex][i].srcStart < translMap[mapIndex][j].srcStart
			})

			mapIndex++
			// fmt.Println("--- End of map\n")
			continue
		}
		// translate
		srcStart, srcEnd, offset := parseRange(input[i])
		item := destination{srcStart, srcEnd, offset}
		// fmt.Println("Src from: ", srcStart, " to: ", srcEnd, " offset ", offset)
		translMap[mapIndex] = append(translMap[mapIndex], item)
	}

	// apply translations
	for i := 0; i < 7; i++ {
		// take seed range
		// sX := seeds[0][0]
		// sY := seeds[0][1]

		// case: it is inside
		// if isInRange((x, translMap[i]))
		// case: first eleme inside range, second outside,not in next range
	}

	fmt.Println("Translation list: ", translMap)

	return res
}

func isInRange(x int, start int, end int) bool {
	return x >= start && x <= end
}

func isBelowLeft(x int, start int) bool {
	return x < start
}

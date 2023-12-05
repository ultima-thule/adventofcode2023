package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Day 04 solution
func fertilizer(filename string, calcFun func([]string) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	var result int
	input := []string{}

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	result = calcFun(input)

	return result
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
	values := make([]int, 0, len(tmp))

	for _, raw := range tmp {
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

// Solve puzzle no 2
func puzzle2(input []string) int {
	var res int = -1

	translated := false
	seedsTmp := splitToInts(input[0])

	var seeds []int

	// mapResults := map[int]int{}

	for i := seedsTmp[0]; i < seedsTmp[0]+seedsTmp[1]; i++ {
		seeds = append(seeds, i)
		// mapResults[i] = 0
	}

	for i := seedsTmp[2]; i < seedsTmp[2]+seedsTmp[3]; i++ {
		seeds = append(seeds, i)
		// mapResults[i] = 0
	}

	fmt.Println("Seeds list: ", seeds)

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

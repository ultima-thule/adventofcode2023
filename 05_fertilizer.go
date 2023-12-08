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

	seeds := [][]int{}

	for x := 0; x < len(seedsTmp); x += 2 {
		tmp1 := []int{seedsTmp[x], seedsTmp[x] + seedsTmp[x+1] - 1}
		seeds = append(seeds, tmp1)
	}
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
		// fmt.Println(input[i], " ", i, " ", len(input))
		if input[i] == "" {
			// end of map
			// fmt.Println(translMap[mapIndex])
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
	for k := 0; k < 7; k++ {
		// sort ranges

		sortSlice(translMap[k])
		translMap[k] = fillInSlice(translMap[k])
		sortSlice(translMap[k])
	}

	// apply translations
	for k := 0; k < 7; k++ {
		fmt.Println("=== Applying map no ", k)
		fmt.Println("= Seeds: ", seeds)

		tmp := make([][]int, 0)

		for i := 0; i < len(seeds); i++ {
			tmp = append(tmp, makeTranslation(seeds[i], translMap[k])...)
		}
		if tmp != nil {
			seeds = append(tmp)
		}
	}

	sortSlice2(seeds)
	fmt.Println("Translated list: ", seeds)

	res = seeds[0][0]

	return res
}

func isInRange(x int, start int, end int) bool {
	return x >= start && x <= end
}

func isBelowLeft(x int, start int) bool {
	return x < start
}

func fillInSlice(sl []destination) []destination {
	// var last int = -1

	if len(sl) < 1 {
		return nil
	}

	if sl[0].srcStart != 0 {
		sl = append(sl, destination{0, sl[0].srcStart - 1, 0})
	}

	return sl
}

func sortSlice(sl []destination) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].srcStart < sl[j].srcStart
	})
}

func sortSlice2(sl [][]int) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i][0] < sl[j][0]
	})
}

func makeTranslation(seeds []int, destRange []destination) [][]int {
	// fmt.Println("Start seeds: ", seeds)
	translated := make([][]int, 0)

	for i := 0; i < len(destRange); i++ {
		oStart, oEnd := getOverlap(seeds[0], seeds[1], destRange[i].srcStart, destRange[i].srcEnd)
		if oStart == -1 && oEnd == -1 {
			continue
		}
		fmt.Println("Overlap: ", oStart, " - ", oEnd)

		item := []int{oStart + destRange[i].offset, oEnd + destRange[i].offset}
		translated = append(translated, item)
	}

	if len(translated) == 0 {
		fmt.Println(seeds, " => ", seeds, "\n")
		translated = append(translated, seeds)
		return translated
	}
	fmt.Println(seeds, " => ", translated, "\n")
	return translated
}

func getOverlap(aStart int, aEnd int, bStart int, bEnd int) (int, int) {
	// fmt.Println("Search for overlap [", aStart, " ", aEnd, "] and [", bStart, " ", bEnd, "]")
	if bStart > aEnd || aStart > bEnd {
		return -1, -1
	}
	oStart := max(aStart, bStart)
	oEnd := min(aEnd, bEnd)

	return oStart, oEnd
}

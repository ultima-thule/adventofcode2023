package main

import (
	"sort"
	"strings"
	"time"
)

type destination struct {
	srcStart int
	srcEnd   int
	offset   int
}

// Solve puzzle no 1
func puzzle05_1(input []string) int {
	defer timeTrack(time.Now(), "part 1")

	var res int = -1

	translated := false
	seeds := splitToInts(input[0])

	// take every seed
	for j := 0; j < len(seeds); j++ {
		translated = false

		// run all translations
		for i := 2; i < len(input); i++ {
			// start of map
			if strings.Contains(input[i], "map") {
				continue
			}
			// end of map
			if input[i] == "" {
				translated = false
				continue
			}
			// translate
			srcStart, srcEnd, offset := parseRange(input[i])
			if !translated && srcStart <= seeds[j] && seeds[j] <= srcEnd {
				seeds[j] += offset
				translated = true
			}
		}
		if res == -1 {
			res = seeds[j]
		}
		// find lowest number
		res = min(res, seeds[j])
	}

	return res
}

// convert into range: start, stop, offset to calculate destination
func parseRange(line string) (int, int, int) {
	split := splitToInts(line)

	if len(split) != 3 {
		return -1, -1, 0
	}
	return split[1], split[1] + split[2] - 1, split[0] - split[1]
}

// Solve puzzle no 2
func puzzle05_2(input []string) int {
	defer timeTrack(time.Now(), "part 2")

	var res int = -1

	seedsTmp := splitToInts(input[0])
	seeds := [][]int{}

	// calculate start and end of seeds ranges
	for x := 0; x < len(seedsTmp); x += 2 {
		tmp1 := []int{seedsTmp[x], seedsTmp[x] + seedsTmp[x+1] - 1}
		seeds = append(seeds, tmp1)
	}

	// read all translations into structure
	translMap := make([][]destination, 7)
	var mapIndex int = 0

	// convert all maps into ranges
	for i := 2; i < len(input); i++ {
		// start of map
		if strings.Contains(input[i], "map") {
			continue
		}
		// end of map
		if input[i] == "" {
			mapIndex++
			continue
		}

		// convert into range
		srcStart, srcEnd, offset := parseRange(input[i])
		item := destination{srcStart, srcEnd, offset}
		translMap[mapIndex] = append(translMap[mapIndex], item)
	}

	// order ranges and add missing fillers with offset=0
	for k := 0; k < 7; k++ {
		sortRanges(translMap[k])
		translMap[k] = fillInSlice(translMap[k])
		sortRanges(translMap[k])
	}

	// apply translations to all seeds
	for k := 0; k < 7; k++ {
		tmp := make([][]int, 0)

		for i := 0; i < len(seeds); i++ {
			tmp = append(tmp, makeTranslation(seeds[i], translMap[k])...)
		}
		if tmp != nil {
			seeds = append(tmp)
		}
	}

	sortSlice(seeds)

	res = seeds[0][0]

	return res
}

// add empty ranges with offset = 0
func fillInSlice(sl []destination) []destination {
	if len(sl) < 1 {
		return nil
	}
	if sl[0].srcStart != 0 {
		sl = append(sl, destination{0, sl[0].srcStart - 1, 0})
	}
	return sl
}

// sort ranges nased on start of range
func sortRanges(sl []destination) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].srcStart < sl[j].srcStart
	})
}

// sort slice of slice based on first elem
func sortSlice(sl [][]int) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i][0] < sl[j][0]
	})
}

// translate seeds positions into new range
func makeTranslation(seeds []int, destRange []destination) [][]int {
	translated := make([][]int, 0)

	for i := 0; i < len(destRange); i++ {
		oStart, oEnd := getOverlap(seeds[0], seeds[1], destRange[i].srcStart, destRange[i].srcEnd)
		// no overlap found
		if oStart == -1 && oEnd == -1 {
			continue
		}
		// calculate destination range
		item := []int{oStart + destRange[i].offset, oEnd + destRange[i].offset}
		translated = append(translated, item)
	}

	// no translation found, return original seeds
	if len(translated) == 0 {
		translated = append(translated, seeds)
	}
	return translated
}

// get overlap of two ranges
func getOverlap(aStart int, aEnd int, bStart int, bEnd int) (int, int) {
	if bStart > aEnd || aStart > bEnd {
		return -1, -1
	}
	oStart := max(aStart, bStart)
	oEnd := min(aEnd, bEnd)

	return oStart, oEnd
}

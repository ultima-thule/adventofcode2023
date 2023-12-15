package main

import (
	"bufio"
	"fmt"
	"sort"
)

func tilt(filename string, calcFunc func(map[string]bool, map[string]bool, int, int) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	roundRocks, cubeRocks, maxRows, maxCols := prepData14(input)
	res += calcFunc(roundRocks, cubeRocks, maxRows, maxCols)

	return res
}

func prepData14(input []string) (map[string]bool, map[string]bool, int, int) {
	roundRocks := make(map[string]bool, 0)
	cubeRocks := make(map[string]bool, 0)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			b := input[i][j]
			key := fmt.Sprintf("%d_%d", i, j)
			switch b {
			case 'O':
				roundRocks[key] = true
			case '#':
				cubeRocks[key] = true
			case '.':
			}
		}
	}

	return roundRocks, cubeRocks, len(input), len(input[0])
}

// Solve puzzle no 14
func puzzle14(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) int {
	// printTilt(roundRocks, cubeRocks, maxRows, maxCols)

	moveNorth(roundRocks, cubeRocks, maxRows, maxCols)
	return countLoad(roundRocks, maxRows, maxCols)
}

// Solve puzzle no 14 part 2
func puzzle14_2(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) int {
	// printTilt(roundRocks, cubeRocks, maxRows, maxCols)

	cycleLen := 0
	it := 0

	visited := make(map[string]int, 0)

	for {
		moveNorth(roundRocks, cubeRocks, maxRows, maxCols)
		moveWest(roundRocks, cubeRocks, maxRows, maxCols)
		moveSouth(roundRocks, cubeRocks, maxRows, maxCols)
		moveEast(roundRocks, cubeRocks, maxRows, maxCols)
		it++

		hash := getHash(roundRocks)

		if visited[hash] != 0 {
			cycleLen = it - visited[hash]
			break
		} else {
			visited[hash] = it
		}
	}

	modulo := (1000000000 - it) % cycleLen

	for i := 0; i < modulo; i++ {
		moveNorth(roundRocks, cubeRocks, maxRows, maxCols)
		moveWest(roundRocks, cubeRocks, maxRows, maxCols)
		moveSouth(roundRocks, cubeRocks, maxRows, maxCols)
		moveEast(roundRocks, cubeRocks, maxRows, maxCols)
	}

	return countLoad(roundRocks, maxRows, maxCols)
}

// generate  unique hash for map
func getHash(m map[string]bool) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	res := ""
	for k := 0; k < len(keys); k++ {
		res = res + keys[k]
	}
	return res
}

// count load on north
func countLoad(roundRocks map[string]bool, maxRows int, maxCols int) int {
	res := 0

	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			key := fmt.Sprintf("%d_%d", r, c)
			if roundRocks[key] {
				res += (maxRows - r)
			}
		}
	}
	return res
}

func moveNorth(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for c := 0; c < maxCols; c++ {
		var lastRock int = -1
		for r := 0; r < maxRows; r++ {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = r
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", lastRock+1, c)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock += 1
				continue
			}
		}
	}
}

func moveSouth(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for c := 0; c < maxCols; c++ {
		var lastRock int = maxRows
		for r := maxRows - 1; r >= 0; r-- {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = r
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", lastRock-1, c)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock -= 1
				continue
			}
		}
	}
}

func moveWest(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		var lastRock int = -1
		for c := 0; c < maxCols; c++ {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = c
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", r, lastRock+1)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock += 1
				continue
			}
		}
	}
}

func moveEast(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		var lastRock int = maxCols
		for c := maxCols; c >= 0; c-- {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = c
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", r, lastRock-1)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock -= 1
				continue
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
)

func tilt(filename string) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	input := make([]string, 0)
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	roundRocks, cubeRocks, maxRows, maxCols := prepData14(input)
	res += puzzle14(roundRocks, cubeRocks, maxRows, maxCols)

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
	printTilt(roundRocks, cubeRocks, maxRows, maxCols)

	// mapNorth := make(map[string]map[string]bool, 0)
	// mapSouth := make(map[string]map[string]bool, 0)
	// mapEast := make(map[string]map[string]bool, 0)
	// mapWest := make(map[string]map[string]bool, 0)

	for i := 0; i < 1000000000; i++ {
		fmt.Println(i)

		// fmt.Println()
		// fmt.Println("NORTH")
		// transformata := mapNorth[fmt.Sprint(roundRocks)]
		// if transformata == nil {
		moveNorth(roundRocks, cubeRocks, maxRows, maxCols)
		// 	mapNorth[tmp] = copy()
		// }
		// else {
		// 	roundRocks = transformata
		// }
		// printTilt(roundRocks, cubeRocks, maxRows, maxCols)
		// fmt.Println()

		// fmt.Println("WEST")
		moveWest(roundRocks, cubeRocks, maxRows, maxCols)
		// mapWest[tmp] = fmt.Sprint(roundRocks)
		// printTilt(roundRocks, cubeRocks, maxRows, maxCols)
		// fmt.Println()

		// fmt.Println("SOUTH")
		moveSouth(roundRocks, cubeRocks, maxRows, maxCols)
		// mapSouth[tmp] = fmt.Sprint(roundRocks)
		// printTilt(roundRocks, cubeRocks, maxRows, maxCols)
		// fmt.Println()

		// fmt.Println("EAST")
		moveEast(roundRocks, cubeRocks, maxRows, maxCols)
		// mapEast[tmp] = fmt.Sprint(roundRocks)
		// fmt.Println()
		// printTilt(roundRocks, cubeRocks, maxRows, maxCols)
	}

	return countLoad(roundRocks, maxRows, maxCols)
}

func compareMaps(map1 map[string]bool, map2 map[string]bool) bool {
	return fmt.Sprint(map1) == fmt.Sprint(map2)
}

func printTilt(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				fmt.Print("#")
				continue
			}
			if roundRocks[key] {
				fmt.Print("O")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

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
		// fmt.Println("-- Column #", c)
		var lastRock int = -1
		for r := 0; r < maxRows; r++ {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = r
				// fmt.Println("Hard rock at", key, "last rock", lastRock)
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", lastRock+1, c)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock += 1
				// fmt.Println("Moving rock from ", key, "to", key2, "last rock", lastRock)
				continue
			}
			// fmt.Println("No rock at ", key, "last rock", lastRock)
		}
	}
}

func moveSouth(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for c := 0; c < maxCols; c++ {
		// fmt.Println("-- Column #", c)
		var lastRock int = maxRows
		for r := maxRows - 1; r >= 0; r-- {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = r
				// fmt.Println("Hard rock at", key, "last rock", lastRock)
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", lastRock-1, c)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock -= 1
				// fmt.Println("Moving rock from ", key, "to", key2, "last rock", lastRock)
				continue
			}
			// fmt.Println("No rock at ", key, "last rock", lastRock)
		}
	}
}

func moveWest(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		// fmt.Println("-- Column #", c)
		var lastRock int = -1
		for c := 0; c < maxCols; c++ {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = c
				// fmt.Println("Hard rock at", key, "last rock", lastRock)
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", r, lastRock+1)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock += 1
				// fmt.Println("Moving rock from ", key, "to", key2, "last rock", lastRock)
				continue
			}
			// fmt.Println("No rock at ", key, "last rock", lastRock)
		}
	}
}

func moveEast(roundRocks map[string]bool, cubeRocks map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		// fmt.Println("-- Column #", c)
		var lastRock int = maxCols
		for c := maxCols; c >= 0; c-- {
			key := fmt.Sprintf("%d_%d", r, c)
			if cubeRocks[key] {
				lastRock = c
				// fmt.Println("Hard rock at", key, "last rock", lastRock)
				continue
			}
			if roundRocks[key] {
				key2 := fmt.Sprintf("%d_%d", r, lastRock-1)
				delete(roundRocks, key)
				roundRocks[key2] = true
				lastRock -= 1
				// fmt.Println("Moving rock from ", key, "to", key2, "last rock", lastRock)
				continue
			}
			// fmt.Println("No rock at ", key, "last rock", lastRock)
		}
	}
}

package main

import (
	"fmt"
)

// Solve puzzle no 1
func puzzle11_1(input [][]byte) int {
	var res int = 0

	foundRows := make(map[int]bool, 0)
	foundCols := make([]int, len(input[0]))
	galaxies := make(map[string]bool, 0)

	addRows := 0

	for i := 0; i < len(input); i++ {
		foundStar := false
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' {
				galaxies[fmt.Sprintf("%d_%d", i+addRows, j)] = true
				foundRows[i+addRows] = true
				foundCols[j] = 1
				foundStar = true
			}
		}
		// expand rows
		if !foundStar {
			addRows++
		}
	}
	maxRows := len(input) + addRows

	// fmt.Println("Galaxies:", galaxies)
	// fmt.Println("Rows:", foundRows)
	// fmt.Println("Cols:", foundCols)

	// printGalaxy(galaxies, 20, 20)

	started := 0
	for i := 0; i < len(foundCols); i++ {
		if foundCols[i] == 1 && started == 0 {
			foundCols[i] = 0
			continue
		}
		if foundCols[i] == 0 {
			started++
			for j := i + 1; j < len(foundCols); j++ {
				if foundCols[j] != 0 {
					foundCols[j] = started
				}
			}
		}
	}
	maxCols := len(input[0]) + started
	// fmt.Println("Cols:", foundCols, "maxCols:", maxCols)

	newGal := make(map[string]bool, 0)
	// fmt.Println("Gals:", galaxies)
	expanded := make([]Point, 0)

	for k := range galaxies {
		var r, c int
		fmt.Sscanf(k, "%d_%d", &r, &c)
		key := fmt.Sprintf("%d_%d", r, c+foundCols[c])
		newGal[key] = true
		expanded = append(expanded, Point{x: r, y: c + foundCols[c]})
	}

	// fmt.Println("Gals:", newGal)
	fmt.Println("Gals:", expanded)

	// for k := range galaxies {
	// 	var r, c int
	// 	fmt.Sscanf(k, "%d_%d", &r, &c)
	// 	galMap = append(galMap, Point{x: r, y: c})
	// }

	fmt.Println()
	fmt.Println("Galaxies:", newGal)
	// fmt.Println("Rows:", foundRows)
	// fmt.Println("Cols:", foundCols)

	printGalaxy(galaxies, len(input), len(input[0]))
	fmt.Println()
	printGalaxy(newGal, maxRows, maxCols)

	// expandCosmos(&grid)

	res, _ = findAllDistances(expanded)

	return res
}

func printGalaxy(input map[string]bool, maxRows int, maxCols int) {
	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxCols; j++ {
			k := fmt.Sprintf("%d_%d", i, j)
			if input[k] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Printf("\n")
	}
}

func findAllDistances(galaxy []Point) (int, map[string]int) {
	distances := make(map[string]int)
	count := 0

	for i := 0; i < len(galaxy); i++ {
		for j := i + 1; j < len(galaxy); j++ {
			dist := distance(galaxy[i], galaxy[j])
			k := fmt.Sprintf("%d_%d", i, j)
			distances[k] = dist
			count += dist
		}
	}

	fmt.Println("Distances:", distances)
	fmt.Println("Count:", count)

	return count, distances
}

func distance(p1 Point, p2 Point) int {
	return abs(p2.y-p1.y) + abs(p2.x-p1.x)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

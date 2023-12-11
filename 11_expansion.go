package main

// Solve puzzle no 1
func puzzle11(input [][]byte, expRate int) int {
	if input == nil {
		return 0
	}

	galaxies, colTransform := expandRows(input, expRate)

	calcTransformation(&colTransform, expRate)
	expanded := expandColumns(galaxies, colTransform)

	return sumDistances(expanded)
}

// read input into point map and expand rows
func expandRows(input [][]byte, expRate int) ([]Point, []int) {
	colTransform := make([]int, len(input[0]))
	points := make([]Point, 0)

	factor := 0

	for i := 0; i < len(input); i++ {
		galaxy := false
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] == '#' {
				points = append(points, Point{x: i + factor, y: j})
				colTransform[j] = 1
				galaxy = true
			}
		}
		if !galaxy {
			factor += (expRate - 1)
		}
	}
	return points, colTransform
}

// calculate column transformation
func calcTransformation(input *[]int, expRate int) {
	factor := 0
	for i := 0; i < len(*input); i++ {
		// skip trailing cols
		if (*input)[i] == 1 && factor == 0 {
			(*input)[i] = 0
			continue
		}
		// empty col found - calculate expansion
		if (*input)[i] == 0 {
			factor += (expRate - 1)
			for j := i + 1; j < len(*input); j++ {
				if (*input)[j] != 0 {
					(*input)[j] = factor
				}
			}
		}
	}
}

// expand columns
func expandColumns(input []Point, transformations []int) []Point {
	expanded := make([]Point, 0)

	for _, p := range input {
		expanded = append(expanded, Point{x: p.x, y: p.y + transformations[p.y]})
	}

	return expanded
}

// sum distances
func sumDistances(galaxy []Point) int {
	count := 0

	for i := 0; i < len(galaxy); i++ {
		for j := i + 1; j < len(galaxy); j++ {
			dist := distance(galaxy[i], galaxy[j])
			count += dist
		}
	}
	return count
}

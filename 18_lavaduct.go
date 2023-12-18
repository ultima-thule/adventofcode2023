package main

import (
	"fmt"
	"sort"
)

type DigPlan struct {
	dir   string
	moves int64
	color string
}

type Point64 struct {
	x int64
	y int64
}

// Solve puzzle no 15 part 1
func puzzle18(input []DigPlan) int64 {
	if input == nil {
		return 0
	}

	visited := make(map[string]bool, 0)
	ranges := make(map[int64][]int64, 0)
	return moveLava(0, 0, &input, visited, ranges)
}

// move based on command r
func moveLava(xS int64, yS int64, input *[]DigPlan, visited map[string]bool, ranges map[int64][]int64) int64 {
	key := fmt.Sprintf("%d_%d", xS, yS)
	visited[key] = true

	polygon := make([]Point64, 0)

	var lastX int64 = xS
	var lastY int64 = yS
	var maxX int64 = 0
	var maxY int64 = 0
	for _, v := range *input {
		x, y := moveOffset18(v.dir, v.moves)

		p := []Point64{{x: lastX, y: lastY}}
		// fmt.Println("Step:", v, "pol", p)
		polygon = append(p, polygon...)

		nextX := lastX + x
		nextY := lastY + y
		maxX = max(maxX, nextX)
		maxY = max(maxY, nextY)

		switch v.dir {
		case "R":
			for j := lastY + 1; j <= nextY; j++ {
				key := fmt.Sprintf("%d_%d", nextX, j)
				visited[key] = true
				addToRange(ranges, nextX, j)
			}
		case "L":
			for j := lastY - 1; j >= nextY; j-- {
				key := fmt.Sprintf("%d_%d", nextX, j)
				visited[key] = true
				addToRange(ranges, nextX, j)
			}
		case "U":
			for i := lastX - 1; i >= nextX; i-- {
				key := fmt.Sprintf("%d_%d", i, nextY)
				visited[key] = true
				addToRange(ranges, i, nextY)
			}
		case "D":
			for i := lastX + 1; i <= nextX; i++ {
				key := fmt.Sprintf("%d_%d", i, nextY)
				visited[key] = true
				addToRange(ranges, i, nextY)
			}
		}
		lastX = nextX
		lastY = nextY
	}

	p := []Point64{{x: 0, y: 0}}
	// fmt.Println("Step:", v, "pol", p)
	polygon = append(p, polygon[:len(polygon)-1]...)

	// printVisited18(visited, maxX+1, maxY+1)
	for _, v := range ranges {
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
		// sort.Ints(v)
	}
	// fmt.Println(polygon)
	// fmt.Println(shoelace(polygon))

	area := shoelace(polygon)
	// fmt.Println("Area:", area)
	interior := pickFormula(area, int64(len(visited)))
	// fmt.Println("Interior:", interior)
	// printVisited18(visited, maxX+1, maxY+1)

	return calcCoverage(int64(len(visited)), interior)
}

// calculate x and y offsets when moving from => to
func moveOffset18(dir string, num int64) (int64, int64) {
	var x, y int64

	switch dir {
	case "R":
		x = 0
		y = num
	case "L":
		x = 0
		y = -num
	case "U":
		x = -num
		y = 0
	case "D":
		x = num
		y = 0
	}

	return x, y
}

func addToRange(ranges map[int64][]int64, x int64, y int64) {
	if ranges[x] == nil {
		ranges[x] = make([]int64, 0)
	}
	ranges[x] = append(ranges[x], y)
}

func shoelace(points []Point64) int64 {
	var sum1, sum2 int64 = 0, 0

	numVert := len(points)
	for i := 0; i < numVert-1; i++ {
		sum1 = sum1 + points[i].x*points[i+1].y
		sum2 = sum2 + points[i].y*points[i+1].x
	}
	sum1 = sum1 + points[numVert-1].x*points[0].y
	sum2 = sum2 + points[0].x*points[numVert-1].y

	area := abs64(sum1-sum2) / 2
	return area
}

func pickFormula(area int64, boundary int64) int64 {
	// area = interior + (1/2 boundary) - 1
	return area + 1 - boundary/2
}

func calcCoverage(boundary int64, interior int64) int64 {
	return boundary + interior
}

package main

import "fmt"

func puzzle21(grid map[string]bool, start Point, maxX int, maxY int) int {
	ret := 0

	visited := map[Point]bool{start: true}

	for i := 0; i < 64; i++ {
		visited = oneStep(grid, visited, maxX, maxY)
	}

	ret = len(visited)

	return ret
}

func oneStep(grid map[string]bool, toVisit map[Point]bool, maxX int, maxY int) map[Point]bool {
	ret := map[Point]bool{}

	for p := range toVisit {
		for _, d := range [4]Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			next := Point{x: p.x + d.x, y: p.y + d.y}
			key := fmt.Sprintf("%d_%d", next.x, next.y)
			if isValid(next, maxX, maxY) && grid[key] == false {
				ret[next] = true
			}
		}
	}

	return ret
}

func isValid(p Point, maxX int, maxY int) bool {
	return p.x >= 0 && p.x < maxX && p.y >= 0 && p.y < maxY
}

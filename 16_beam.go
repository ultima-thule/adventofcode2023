package main

import (
	"fmt"
)

type (
	Direction int
)

const (
	N Direction = iota
	E
	S
	W
)

type Cell struct {
	x int
	y int
	e Direction // direction of entrance to the cell
}

// Solve puzzle no 15 part 1
func puzzle16(input [][]byte) int {
	if input == nil {
		return 0
	}

	visited := make(map[string]bool, 0)

	m := Cell{e: W, x: 0, y: 0}
	move(m, input[0][0], &input, visited)

	return calcVisited(visited, len(input), len(input[0]))
}

// Solve puzzle no 15 part 1
func puzzle16_2(input [][]byte) int {
	if input == nil {
		return 0
	}

	maxRows := len(input)
	maxCols := len(input[0])

	maxRes := 0
	entryPoints := make([]Cell, 0)

	// top row from N, bottom row from S
	for i := 0; i < maxCols; i++ {
		m1 := Cell{e: N, x: 0, y: i}
		m2 := Cell{e: S, x: maxRows - 1, y: i}
		entryPoints = append(entryPoints, m1, m2)
	}
	// left column from W, right column from E
	for i := 0; i < maxRows; i++ {
		m1 := Cell{e: W, x: i, y: 0}
		m2 := Cell{e: E, x: i, y: maxCols - 1}
		entryPoints = append(entryPoints, m1, m2)
	}
	// find maximum beam coverage
	for _, m := range entryPoints {
		visited := make(map[string]bool, 0)
		move(m, input[m.x][m.y], &input, visited)
		maxRes = max(maxRes, calcVisited(visited, maxRows, maxCols))
	}

	return maxRes
}

// move based on command r
func move(m Cell, r byte, input *[][]byte, visited map[string]bool) {
	key := fmt.Sprintf("%d_%d_%s", m.x, m.y, m.e)
	if visited[key] {
		return
	}
	visited[key] = true

	dir := getNextDir(r, m.e)

	for _, v := range dir {
		n := moveOffset(m.e, v)
		nextX := m.x + n.x
		nextY := m.y + n.y

		if nextX >= 0 && nextX < len(*input) && nextY >= 0 && nextY < len((*input)[0]) {
			p := Cell{e: v, x: nextX, y: nextY}
			move(p, (*input)[nextX][nextY], input, visited)
		}
	}
}

// get direction from which next cell will be entered
func getNextDir(r byte, d Direction) []Direction {
	dirs := []map[byte][]Direction{
		{
			'|':  {N},
			'-':  {W, E},
			'\\': {W},
			'/':  {E},
			'.':  {N},
		},
		{
			'|':  {S, N},
			'-':  {E},
			'\\': {S},
			'/':  {N},
			'.':  {E},
		},
		{
			'|':  {S},
			'-':  {W, E},
			'\\': {E},
			'/':  {W},
			'.':  {S},
		},
		{
			'|':  {S, N},
			'-':  {W},
			'\\': {N},
			'/':  {S},
			'.':  {W},
		},
	}

	return dirs[d][r]
}

// calculate x and y offsets when moving from => to
func moveOffset(from Direction, to Direction) Point {
	// N = 0, E = 1, S = 2, W = 3
	res := map[string]Point{
		"0_0": {1, 0},  // N=>N
		"0_1": {0, -1}, // N=>E
		//"0_2": {0, 0}, // N=>S non existing
		"0_3": {0, 1},  // N=> W
		"1_0": {1, 0},  // E=>N
		"1_1": {0, -1}, // E=>E
		"1_2": {-1, 0}, // E=>S
		//"1_3": {0, 0}, // E=>W non existing
		//"2_0": {1, 0},  // S=>N non existing
		"2_1": {0, -1}, // S=>E
		"2_2": {-1, 0}, // S=>S
		"2_3": {0, 1},  // S=>W
		"3_0": {1, 0},  // W=>N
		//"3_1": {0, -1}, // W=>E non existing
		"3_2": {-1, 0}, // W=>S
		"3_3": {0, 1},  // W=>W
	}
	return res[fmt.Sprintf("%d_%d", from, to)]
}

func calcVisited(visited map[string]bool, maxRows int, maxCols int) int {
	cnt := 0

	// res := make(map[string]bool, 0)
	// var r, c int
	// var s string
	// for k, v := range visited {
	// 	fmt.Sprintf("%d_%d_%s", r, c, s)
	// }

	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			keyN := fmt.Sprintf("%d_%d_N", r, c)
			keyS := fmt.Sprintf("%d_%d_S", r, c)
			keyW := fmt.Sprintf("%d_%d_W", r, c)
			keyE := fmt.Sprintf("%d_%d_E", r, c)
			if visited[keyN] || visited[keyS] || visited[keyW] || visited[keyE] {
				cnt++
			}
		}
	}
	return cnt
}

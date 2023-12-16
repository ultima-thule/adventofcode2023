package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type (
	Direction int
	Command   int
)

const (
	N Direction = iota
	E
	S
	W
)

type Mirror struct {
	x       int
	y       int
	entered Direction
}

// Day 16 solution
func beam(filename string, calcFun func([][]byte) int) int {
	// Read the contents of the file
	fmt.Println("=> DataSet: ", filename)

	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	grid := parseInputIntoBytes(content)

	res := calcFun(grid)

	defer timeTrack(time.Now(), "beam")
	fmt.Println()

	return res
}

// Solve puzzle no 15 part 1
func puzzle16(input [][]byte) int {
	if input == nil {
		return 0
	}

	res := 0

	// fmt.Println(input)
	visited := make(map[string]bool, 0)

	m := Mirror{entered: W, x: 0, y: 0}

	move(m, input[0][0], &input, visited)

	res = calcVisited(visited, len(input), len(input[0]))

	return res
}

// Solve puzzle no 15 part 1
func puzzle16_2(input [][]byte) int {
	if input == nil {
		return 0
	}

	maxRows := len(input)
	maxCols := len(input[0])

	maxRes := 0

	// top row from N
	for i := 0; i < maxCols; i++ {
		res := 0
		visited := make(map[string]bool, 0)
		m := Mirror{entered: N, x: 0, y: i}
		move(m, input[0][i], &input, visited)
		res = calcVisited(visited, maxRows, maxCols)
		maxRes = max(res, maxRes)
	}
	// bottom row from S
	for i := 0; i < maxCols; i++ {
		res := 0
		visited := make(map[string]bool, 0)
		m := Mirror{entered: S, x: maxRows - 1, y: i}
		move(m, input[maxRows-1][i], &input, visited)
		res = calcVisited(visited, maxRows, maxCols)
		maxRes = max(res, maxRes)
	}
	// left column from W
	for i := 0; i < maxRows; i++ {
		res := 0
		visited := make(map[string]bool, 0)
		m := Mirror{entered: W, x: i, y: 0}
		move(m, input[i][0], &input, visited)
		res = calcVisited(visited, maxRows, maxCols)
		maxRes = max(res, maxRes)
	}
	// right column from E
	for i := 0; i < maxRows; i++ {
		res := 0
		visited := make(map[string]bool, 0)
		m := Mirror{entered: E, x: i, y: maxCols - 1}
		move(m, input[i][maxCols-1], &input, visited)
		res = calcVisited(visited, maxRows, maxCols)
		maxRes = max(res, maxRes)
	}

	return maxRes
}

func m(m Mirror) {
}

func move(m Mirror, r byte, input *[][]byte, visited map[string]bool) {
	// mark as visited
	key := fmt.Sprintf("%d_%d_%s", m.x, m.y, m.entered)
	// fmt.Println("\nVisiting ", key, "from direction", m.entered.String(), "char", string(r))
	if visited[key] {
		return
	}

	visited[key] = true

	// printVisited(visited, len(*input), len((*input)[0]))

	var dir []Direction
	switch m.entered {
	case N:
		// fmt.Println("Entering from N")
		dir = fromNorth(r)
	case S:
		// fmt.Println("Entering from S")
		dir = fromSouth(r)
	case E:
		// fmt.Println("Entering from E")
		dir = fromEast(r)
	case W:
		// fmt.Println("Entering from W")
		dir = fromWest(r)
	}
	//	fmt.Println("Found", len(dir), "options")

	for _, v := range dir {

		// fmt.Println("Option", k)
		// fmt.Println("Rune ", r, " => dir", dir)

		n := nextMove(m.entered, v)
		nextX := m.x + n.x
		nextY := m.y + n.y

		if nextX >= 0 && nextX < len(*input) && nextY >= 0 && nextY < len((*input)[0]) {
			// fmt.Println("Next move: ", m.entered, "=>", v, "cmd", string(r), "next point: ", nextX, nextY)
			p := Mirror{entered: v, x: nextX, y: nextY}
			move(p, (*input)[nextX][nextY], input, visited)
		} else {
			// fmt.Println("-Deadlock move: ", m.entered, "=>", v, "next point: ", nextX, nextY)
		}
	}
}

func fromNorth(r byte) []Direction {
	if r == '|' {
		return []Direction{N}
	}
	if r == '-' {
		return []Direction{W, E}
	}
	if r == '\\' {
		return []Direction{W}
	}
	if r == '/' {
		return []Direction{E}
	}
	return []Direction{N}
}

func fromSouth(r byte) []Direction {
	if r == '|' {
		return []Direction{S}
	}
	if r == '-' {
		return []Direction{W, E}
	}
	if r == '\\' {
		return []Direction{E}
	}
	if r == '/' {
		return []Direction{W}
	}
	return []Direction{S}
}

func fromWest(r byte) []Direction {
	if r == '-' {
		return []Direction{W}
	}
	if r == '|' {
		return []Direction{S, N}
	}
	if r == '\\' {
		return []Direction{N}
	}
	if r == '/' {
		return []Direction{S}
	}
	return []Direction{W}
}

func fromEast(r byte) []Direction {
	if r == '-' {
		return []Direction{E}
	}
	if r == '|' {
		return []Direction{S, N}
	}
	if r == '\\' {
		return []Direction{S}
	}
	if r == '/' {
		return []Direction{N}
	}
	return []Direction{E}
}

func nextMove(prev Direction, next Direction) Point {
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
	key := fmt.Sprintf("%d_%d", prev, next)
	// fmt.Println("Mapping:", prev.String(), "=>", next.String(), "(", key, ") to", res[key])

	return res[key]
}

func (d Direction) String() string {
	return []string{"N", "E", "S", "W"}[d]
}

func printVisited(visited map[string]bool, maxRows int, maxCols int) {
	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			keyN := fmt.Sprintf("%d_%d_N", r, c)
			keyS := fmt.Sprintf("%d_%d_S", r, c)
			keyW := fmt.Sprintf("%d_%d_W", r, c)
			keyE := fmt.Sprintf("%d_%d_E", r, c)
			if visited[keyN] || visited[keyS] || visited[keyW] || visited[keyE] {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

func calcVisited(visited map[string]bool, maxRows int, maxCols int) int {
	cnt := 0
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

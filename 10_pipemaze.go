package main

import "fmt"

type Point struct {
	x int
	y int
}

// Solve puzzle no 1
func puzzle10_1(input []string, start Point) int {
	fmt.Println("Start:", start)
	// fmt.Println("Input:", input)

	points := findStartConnections(start, input)

	if len(points) != 2 {
		return 0
	}

	// fmt.Println("Possible start connections: ", points, "\n")
	visited := make(map[string]bool, 0)
	conc := fmt.Sprintf("%d_%d", points[0].x, points[0].y)
	visited[conc] = true
	conc = fmt.Sprintf("%d_%d", start.x, start.y)
	visited[conc] = true

	followMaze(points[0], start, visited, input)

	return len(visited) / 2
}

// Solve puzzle no 2
// func puzzle10_2(input []string) int {
// 	if input == nil {
// 		return 0
// 	}

// 	res := 0
// 	steps := make([]int, 0)

// 	_, s := reduce(input, steps, true)
// 	for i := len(s) - 1; i >= 0; i-- {
// 		res = s[i] - res
// 	}

// 	return res
// }

func findStartConnections(start Point, input []string) []Point {
	res := make([]Point, 0)

	if start.x > 0 {
		top := Point{x: start.x - 1, y: start.y}
		// fmt.Println("Top pos:", top)
		if checkConnection(input[top.x][top.y], "t") {
			res = append(res, top)
		}
	}

	if start.y > 0 {
		left := Point{x: start.x, y: start.y - 1}
		// fmt.Println("Left pos:", left)
		if checkConnection(input[left.x][left.y], "l") {
			res = append(res, left)
		}
	}

	if start.x < len(input[0]) {
		down := Point{x: start.x + 1, y: start.y}
		// fmt.Println("Down pos:", down)
		if checkConnection(input[down.x][down.y], "d") {
			res = append(res, down)
		}
	}

	if start.y < len(input) {
		right := Point{x: start.x, y: start.y + 1}
		// fmt.Println("Right pos:", right)
		if checkConnection(input[right.x][right.y], "r") {
			res = append(res, right)
		}
	}

	// fmt.Println("Possible connections", res)

	return res
}

func checkConnection(b byte, t string) bool {
	s := string(b)

	// fmt.Println("Check: ", s, t)

	switch t {
	case "l":
		return s == "-" || s == "L" || s == "F"
	case "t":
		return s == "|" || s == "7" || s == "F"
	case "r":
		return s == "-" || s == "7" || s == "J"
	case "d":
		return s == "|" || s == "L" || s == "J"
	}
	return false
}

func followMaze(start Point, end Point, visited map[string]bool, input []string) {
	// fmt.Println("\n- Entered ", start)
	// reached the start point
	if start.x == end.x && start.y == end.y {
		return
	}

	// fmt.Println("Follow ", start, "visited: ", visited)

	conn := findConnections(start, input)
	for i := 0; i < len(conn); i++ {
		conc := fmt.Sprintf("%d_%d", conn[i].x, conn[i].y)
		if !visited[conc] {
			// fmt.Println("Not visited: ", conc, "\n")
			// fmt.Println("=> ", conn[i].x, ",", conn[i].y)
			visited[conc] = true
			followMaze(conn[i], end, visited, input)
		}
	}
}

func findConnections(start Point, input []string) []Point {
	res := make([]Point, 0)

	curr := string(input[start.x][start.y])

	if start.x > 0 && (curr == "|" || curr == "L" || curr == "J") {
		top := Point{x: start.x - 1, y: start.y}
		// fmt.Println("Top pos:", top)
		if checkConnection2(curr, input[top.x][top.y], "t") {
			res = append(res, top)
		}
	}

	if start.y > 0 && (curr == "-" || curr == "J" || curr == "7") {
		left := Point{x: start.x, y: start.y - 1}
		// fmt.Println("Left pos:", left)
		if checkConnection2(curr, input[left.x][left.y], "l") {
			res = append(res, left)
		}
	}

	if start.x < len(input[0]) && (curr == "|" || curr == "7" || curr == "F") {
		down := Point{x: start.x + 1, y: start.y}
		// fmt.Println("Down pos:", down)
		if checkConnection2(curr, input[down.x][down.y], "d") {
			res = append(res, down)
		}
	}

	if start.y < len(input) && (curr == "-" || curr == "L" || curr == "F") {
		right := Point{x: start.x, y: start.y + 1}
		// fmt.Println("Right pos:", right)
		if checkConnection2(curr, input[right.x][right.y], "r") {
			res = append(res, right)
		}
	}

	// fmt.Println("Possible connections", res)

	return res
}

func checkConnection2(curr string, b byte, t string) bool {
	s := string(b)

	// fmt.Println("Check: ", s, t)

	switch t {
	case "l":
		return s == "-" || s == "L" || s == "F"
	case "t":
		return s == "|" || s == "7" || s == "F"
	case "r":
		return s == "-" || s == "7" || s == "J"
	case "d":
		return s == "|" || s == "L" || s == "J"
	}
	return false
}

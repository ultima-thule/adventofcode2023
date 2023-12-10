package main

import "fmt"

type Point struct {
	x int
	y int
}

// Solve puzzle no 1
func puzzle10_1(input []string, start Point) int {
	// fmt.Println("Start:", start)

	points := findStartConnections(start, input)

	if len(points) != 2 {
		return 0
	}

	visited := make(map[string]bool, 0)
	conc := fmt.Sprintf("%d_%d", points[0].x, points[0].y)
	visited[conc] = true
	conc = fmt.Sprintf("%d_%d", start.x, start.y)
	visited[conc] = true

	corners := make([]Point, 0)
	corners = append(corners, Point{x: start.x, y: start.y})
	corners = append(corners, Point{x: points[0].x, y: points[0].y})

	followMaze(points[0], start, visited, &corners, input)

	corners = append(corners, Point{x: start.x, y: start.y})

	return len(visited) / 2
}

// Solve puzzle no 2
func puzzle10_2(input []string, start Point) int {
	// fmt.Println("Start:", start)

	points := findStartConnections(start, input)
	// fmt.Println("Start conns: ", points)

	if len(points) != 2 {
		return 0
	}

	visited := make(map[string]bool, 0)
	corners := make([]Point, 0)
	conc := fmt.Sprintf("%d_%d", points[0].x, points[0].y)
	visited[conc] = true
	conc = fmt.Sprintf("%d_%d", start.x, start.y)
	visited[conc] = true

	corners = append(corners, Point{x: start.x, y: start.y})
	corners = append(corners, Point{x: points[0].x, y: points[0].y})

	followMaze(points[0], start, visited, &corners, input)

	corners = append(corners, Point{x: start.x, y: start.y})

	// fmt.Println("Visited:", visited)
	// fmt.Println("Corners: ", corners)

	output := drawX(input, visited)
	fmt.Println()

	fmt.Println(output[0][0])

	// res := calcCrosses(input, output)
	// output = drawO(output, visited)
	res := drawLR(output, corners, visited)

	// inner := make(map[string]bool, 0)
	// fillInner(Point{start.x + 1, start.y + 1}, output, inner)

	// fmt.Println("Puzzle result: ", inner)

	return res
}

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

	if start.x < len(input) {
		down := Point{x: start.x + 1, y: start.y}
		// fmt.Println("Down pos:", down)
		if checkConnection(input[down.x][down.y], "d") {
			res = append(res, down)
		}
	}

	if start.y < len(input[0]) {
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

func followMaze(start Point, end Point, visited map[string]bool, corners *[]Point, input []string) {
	// fmt.Println("\n- Entered ", start)
	// reached the start point
	if start.x == end.x && start.y == end.y {
		return
	}

	// fmt.Println("Follow ", start, "visited: ", visited)

	conn := findConnections(start, input)
	// fmt.Println("Found:", conn)
	for i := 0; i < len(conn); i++ {
		p := fmt.Sprintf("%d_%d", conn[i].x, conn[i].y)
		if !visited[p] {
			// fmt.Println("Following:", p)
			// fmt.Println("Not visited: ", conc, "\n")
			// fmt.Println("=> ", conn[i].x, ",", conn[i].y)
			visited[p] = true
			*corners = append(*corners, Point{x: conn[i].x, y: conn[i].y})

			followMaze(conn[i], end, visited, corners, input)
		}
	}
}

func findConnections(start Point, input []string) []Point {
	// fmt.Println("At point ", start)
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

	if start.y < len(input[0]) && (curr == "-" || curr == "L" || curr == "F") {
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

func drawX(input []string, borderMap map[string]bool) []string {
	output := make([]string, 0)

	for i := 0; i < len(input); i++ {
		row := ""

		for j := 0; j < len(input[i]); j++ {
			conc := fmt.Sprintf("%d_%d", i, j)
			if borderMap[conc] {
				row += "X"
			} else {
				// row += string(input[i][j])
				row += string(".")
			}
		}
		fmt.Println("Row:", row)
		output = append(output, row)
	}

	return output
}

func drawO(input []string, borderMap map[string]bool) []string {
	output := make([]string, 0)

	for i := 0; i < len(input); i++ {
		row := ""
		inside := false

		for j := 0; j < len(input[i]); j++ {

			curr := string(input[i][j])
			// left border
			if curr == "X" && i == 0 {
				// fmt.Println("Left X border")
				inside = true
				row += curr
				continue
			}
			// not the last column
			if j != len(input[i])-1 {
				right := string(input[i][j+1])
				if curr == "X" && inside && right != "X" {
					// fmt.Println("I am ", curr, " and inside and not X (", right, " to right")
					inside = false
					row += curr
					continue
				}
				if right == "X" && !inside {
					// fmt.Println("I am ", curr, " and not inside and ", right, " to right")
					inside = true
					row += "O"
					continue
				}
			} else {
				prev := string(input[i][j-1])
				if curr != "X" && prev == "X" {
					row += "O"
					continue
				}
			}
			// just append
			// fmt.Println("I'm other: ", curr)
			row += curr
		}
		fmt.Println("Row:", row)
		output = append(output, row)
	}

	return output
}

func drawLR(input []string, corners []Point, borderMap map[string]bool) int {
	leftPoints := make(map[string]bool, 0)
	rightPoints := make(map[string]bool, 0)

	maxX := len(input)
	maxY := len(input[0])

	// fmt.Println("Corners:", corners)
	// fmt.Println("Max :", maxX, ",", maxY)

	// fmt.Println("check :", string(input[0][4]))

	// prevMove := ""
	// nextMove := ""

	isFirst := true

	for i := 0; i < len(corners)-1; i++ {
		p := corners[i]
		n := corners[i+1]

		// fmt.Println("Point:", p, " next: ", n)

		markLeft := make([]Point, 3)
		markRight := make([]Point, 3)

		// prevMove = nextMove

		if n.y == p.y+1 {
			// move right horizontally
			// nextMove = "R"
			// fmt.Println("=>")
			markLeft[0] = Point{p.x - 1, p.y - 1}
			markLeft[1] = Point{p.x - 1, p.y}
			markLeft[2] = Point{p.x - 1, p.y + 1}

			markRight[0] = Point{p.x + 1, p.y - 1}
			markRight[1] = Point{p.x + 1, p.y}
			markRight[2] = Point{p.x + 1, p.y + 1}

			// fmt.Println("L:", left, "R:", right)
		} else if n.x == p.x+1 {
			// move down vertically OK
			// nextMove = "D"
			// fmt.Println("V")

			markLeft[0] = Point{p.x, p.y + 1}
			markLeft[1] = Point{p.x + 1, p.y + 1}

			markRight[0] = Point{p.x - 1, p.y - 1}
			markRight[1] = Point{p.x, p.y - 1}
			if !isFirst {
				markLeft[2] = Point{p.x - 1, p.y + 1}
				markRight[2] = Point{p.x + 1, p.y - 1}
				isFirst = false

			}

		} else if n.y == p.y-1 {
			// move left horizontally
			// nextMove = "L"
			// fmt.Println("<=")

			markLeft[0] = Point{p.x + 1, p.y - 1}
			markLeft[1] = Point{p.x + 1, p.y}
			markLeft[2] = Point{p.x + 1, p.y + 1}

			markRight[0] = Point{p.x - 1, p.y - 1}
			markRight[1] = Point{p.x - 1, p.y}
			markRight[2] = Point{p.x - 1, p.y + 1}

			// fmt.Println("L:", left, "R:", right)
		} else if n.x == p.x-1 {
			// move up vertically OK
			// nextMove = "U"
			// fmt.Println("I")
			// fmt.Println("L:", left, "R:", right)
			markLeft[0] = Point{p.x - 1, p.y - 1}
			markLeft[1] = Point{p.x, p.y - 1}
			markLeft[2] = Point{p.x + 1, p.y - 1}

			markRight[0] = Point{p.x - 1, p.y + 1}
			markRight[1] = Point{p.x, p.y + 1}
			markRight[2] = Point{p.x + 1, p.y + 1}
		}

		// fmt.Println("-- left: ", left)
		// fmt.Println("-- right: ", right)

		for i := 0; i < len(markRight); i++ {
			// fmt.Println("RP:", markRight[i])
			if conc := isValidPoint(markRight[i], maxX, maxY); conc != "" {
				// if right.x >= 0 && right.x < len(input) && right.y >= 0 && right.y < len(input[0]) {
				if borderMap[conc] == false && leftPoints[conc] == false {
					rightPoints[conc] = true
				}
			}
		}
		for i := 0; i < len(markLeft); i++ {
			// fmt.Println("LP:", markLeft[i])
			if conc := isValidPoint(markLeft[i], maxX, maxY); conc != "" {
				// if left.x >= 0 && left.x < len(input) && left.y >= 0 && left.y < len(input[0]) {
				if borderMap[conc] == false && rightPoints[conc] == false {
					leftPoints[conc] = true
				}
			}
		}

	}

	resL, resR := calcInside(input, leftPoints, rightPoints)
	res := min(resL, resR)

	return res
}

func isValidPoint(p Point, maxX int, maxY int) string {
	if p.x >= 0 && p.x < maxX && p.y >= 0 && p.y < maxY {
		return fmt.Sprintf("%d_%d", p.x, p.y)
	}
	return ""
}

func calcInner(input []string) int {
	res := 0
	found := 0
	inside := false

	for i := 0; i < len(input); i++ {
		inside = false
		found = 0
		// fmt.Println(input[i])
		for j := 0; j < len(input[i]); j++ {
			curr := string(input[i][j])
			if curr == "O" {
				found++
				inside = true
				continue
			}
			if inside && curr != "X" && curr != "R" {
				found++
				continue
			}
			if inside && (curr == "X" || curr == "R") {
				inside = false
				continue
			}

		}
		res += found
		// fmt.Println("found:", found)
	}

	return res
}

func fillInner(start Point, input []string, inner map[string]bool) {
	toFill := findToFill(start, input, inner)
	if len(toFill) == 0 {
		return
	}
	for i := 0; i < len(toFill); i++ {
		conc := fmt.Sprintf("%d_%d", toFill[i].x, toFill[i].y)
		inner[conc] = true
		fillInner(toFill[i], input, inner)
	}
}

func findToFill(start Point, input []string, inner map[string]bool) []Point {
	ret := make([]Point, 0)
	p1 := Point{start.x, start.y - 1}
	p2 := Point{start.x, start.y + 1}
	p3 := Point{start.x - 1, start.y}
	p4 := Point{start.x + 1, start.y}

	if conc := isValidPoint(p1, len(input), len(input[0])); conc != "" && input[p1.x][p1.y] != 'X' && input[p1.x][p1.y] != 'I' && !inner[conc] {
		ret = append(ret, p1)
	}
	if conc := isValidPoint(p2, len(input), len(input[0])); conc != "" && input[p2.x][p2.y] != 'X' && input[p2.x][p2.y] != 'I' && !inner[conc] {
		ret = append(ret, p2)
	}
	if conc := isValidPoint(p3, len(input), len(input[0])); conc != "" && input[p3.x][p3.y] != 'X' && input[p3.x][p3.y] != 'I' && !inner[conc] {
		ret = append(ret, p3)
	}
	if conc := isValidPoint(p4, len(input), len(input[0])); conc != "" && input[p4.x][p4.y] != 'X' && input[p4.x][p4.y] != 'I' && !inner[conc] {
		ret = append(ret, p4)
	}

	return ret
}

func calcCrosses(original []string, input []string) int {
	insidePoints := 0
	prevOrig := ""
	currOrig := ""
	collapse := false

	for i := 0; i < len(input); i++ {
		fmt.Println("Row: ", original[i])
		fmt.Println("Row: ", input[i])
		currIntersect := 0
		for j := 0; j < len(input[i]); j++ {
			collapse = false
			if string(original[i][j]) == "-" {
				fmt.Println("Border found: ", string(original[i][j]), " intersect counter: ", currIntersect)
				continue
			}
			curr := string(input[i][j])
			prevOrig = currOrig
			currOrig = string(original[i][j])

			if prevOrig == "|" && currOrig == "|" {
				collapse = false
			}

			// fmt.Println("Original: ", currOrig)

			if curr == "X" {

				if !collapse {
					currIntersect++
				} else {
					currIntersect--
				}
				// if currOrig == "L" || currOrig == "F" || currOrig == "S" || currOrig == "|" || (currOrig == "7" && prevOrig != "F" && prevOrig != "L" && prevOrig != "-") || (currOrig == "J" && prevOrig != "F" && prevOrig != "L" && prevOrig != "-") {
				// 	currIntersect++
				// }
				fmt.Println("Border found: ", currOrig, "(", prevOrig, ") intersect counter: ", currIntersect)
				if !collapse && prevOrig == "|" && currOrig == "L" {
					collapse = true
					fmt.Println("Collapsing border: ", currOrig, " intersect counter: ", currIntersect)

				}
				if !collapse && prevOrig == "|" && currOrig == "F" {
					collapse = true
					fmt.Println("Collapsing border: ", currOrig, " intersect counter: ", currIntersect)

				}
				if !collapse && prevOrig == "J" && currOrig == "|" {
					collapse = true
					fmt.Println("Collapsing border: ", currOrig, " intersect counter: ", currIntersect)

				}
				if !collapse && prevOrig == "7" && currOrig == "|" {
					collapse = true
					fmt.Println("Collapsing border: ", currOrig, " intersect counter: ", currIntersect)
				}
				continue
			}

			if currIntersect%2 != 0 {
				fmt.Println("Inside point found: ", curr, " intersect counter: ", currIntersect)
				insidePoints++
				continue
			}
			if currIntersect%2 == 0 {
				fmt.Println("Outside point found: ", curr, " intersect counter: ", currIntersect)
				continue
			}
		}

		fmt.Println("Points: ", insidePoints, "\n")
	}
	return insidePoints
}

func calcCrosses2(original []string, input []string) int {
	insidePoints := 0

	for i := 0; i < len(input); i++ {
		fmt.Println("Row: ", original[i])
		fmt.Println("Row: ", input[i])
		currIntersect := 0
		for j := 0; j < len(input[i]); j++ {
			curr := string(input[i][j])
			currOrig := string(original[i][j])

			if curr == "X" {
				if currOrig != "-" {
					currIntersect++
					fmt.Println("Vertex point found: ", currOrig, " intersect counter: ", currIntersect)
				} else {
					fmt.Println("Border point found: ", currOrig, " intersect counter: ", currIntersect)
				}
				continue
			}

			if currIntersect%2 != 0 {
				fmt.Println("Inside point found: ", currOrig, " intersect counter: ", currIntersect)
				insidePoints++
				continue
			}
			if currIntersect%2 == 0 {
				fmt.Println("Outside point found: ", currOrig, " intersect counter: ", currIntersect)
				continue
			}
		}

		fmt.Println("Points: ", insidePoints, "\n")
	}
	return insidePoints
}

func calcInside(input []string, pointsLeft map[string]bool, pointsRight map[string]bool) (int, int) {
	countL := 0
	countR := 0

	for i := 0; i < len(input); i++ {
		row := ""

		for j := 0; j < len(input[i]); j++ {
			curr := string(input[i][j])
			conc := fmt.Sprintf("%d_%d", i, j)

			if pointsLeft[conc] {
				row += "O"
				continue
			}
			if pointsRight[conc] {
				row += "+"
				continue
			}
			row += curr
		}
		// fmt.Println(row)
		countL += convertRow(row, "O")
		// countR += convertRow(row, "+")
	}

	fmt.Println()

	fmt.Println("Count:", countL, countR)
	fmt.Println()
	return countL, countR
}

func convertRow(input string, letter string) int {
	// fmt.Println("TODO:", input)

	count := 0
	row := ""
	inside := false

	for i := 0; i < len(input)-1; i++ {
		curr := string(input[i])
		// fmt.Println("Item: ", curr, " inside: ", inside)
		if string(input[i]) == letter {
			if string(input[i+1]) == "." {
				inside = true
			}
			count++
			row += letter
			continue
		}
		if string(input[i]) == "." && inside {
			row += letter
			count++
			continue
		}
		row += curr
	}

	fmt.Println(row)
	return count
}

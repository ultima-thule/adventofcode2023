package main

type DigPlan struct {
	dir   string
	moves int64
}

type Point64 struct {
	x int64
	y int64
}

// Solve puzzle no 18 part 1 & 2
func puzzle18(input []DigPlan) int64 {
	if input == nil {
		return 0
	}

	return moveLava(0, 0, &input)
}

// move based on dig plan
func moveLava(xS int64, yS int64, input *[]DigPlan) int64 {
	var boundaryCnt int64 = 0

	polygon := make([]Point64, 0)

	var lastX int64 = xS
	var lastY int64 = yS
	for _, v := range *input {
		vector := moveOffset18(v.dir, v.moves)

		p := []Point64{{x: lastX, y: lastY}}
		polygon = append(p, polygon...)

		nextX := lastX + vector.x
		nextY := lastY + vector.y

		toAdd := map[string]int64{
			"R": nextY - lastY,
			"L": lastY - nextY,
			"U": lastX - nextX,
			"D": nextX - lastX,
		}

		boundaryCnt += toAdd[v.dir]
		lastX = nextX
		lastY = nextY
	}

	// move starting point to the beginning
	polygon = append(polygon[len(polygon)-1:], polygon[:len(polygon)-1]...)

	// calc total area
	area := shoelace(polygon)
	// calc interior points
	interior := pickFormula(area, boundaryCnt)

	return boundaryCnt + interior
}

// calculate x and y offsets when moving from => to
func moveOffset18(dir string, num int64) Point64 {
	v := Point64{0, 0}

	switch dir {
	case "R":
		v.y = num
	case "L":
		v.y = -num
	case "U":
		v.x = -num
	case "D":
		v.x = num
	}
	return v
}

// sholace formula for calculating area
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

// pick formula for calculating interior
func pickFormula(area int64, boundary int64) int64 {
	// area = interior + (1/2 boundary) - 1
	return area + 1 - boundary/2
}

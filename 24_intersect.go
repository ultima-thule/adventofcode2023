package main

type Vector struct {
	x  float64
	y  float64
	z  float64
	vx float64
	vy float64
	vz float64
}

// Solve puzzle no 24 part 1
func puzzle24(vec []Vector, minVal float64, maxVal float64) int {
	if vec == nil {
		return 0
	}

	res := 0

	for i := 0; i < len(vec); i++ {
		for j := i + 1; j < len(vec); j++ {
			px, py, ok := calcIntersection(vec[i], vec[j])
			if ok && within(px, minVal, maxVal) && within(py, minVal, maxVal) {
				res++
			}
		}
	}

	return res
}

func calcIntersection(a Vector, b Vector) (float64, float64, bool) {
	dx := b.x - a.x
	dy := b.y - a.y

	det := a.vx*b.vy - a.vy*b.vx

	if det != 0 {
		u := float64((dx*b.vy - dy*b.vx) / det)
		v := float64((dx*a.vy - dy*a.vx) / det)
		if u >= 0 && v >= 0 {
			return float64(a.x + a.vx*u), float64(a.y + a.vy*u), true
		}
	}

	return 0, 0, false
}

func within(p float64, minVal float64, maxVal float64) bool {
	return p >= minVal && p <= maxVal
}

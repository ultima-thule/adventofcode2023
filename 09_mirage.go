package main

// Solve puzzle no 1
func puzzle09_1(input []int) int {
	if input == nil {
		return 0
	}

	res := input[len(input)-1]
	steps := make([]int, 0)

	_, n := reduce(input, steps, false)
	for i := 0; i < len(n); i++ {
		res += n[i]
	}

	return res
}

// Solve puzzle no 2
func puzzle09_2(input []int) int {
	if input == nil {
		return 0
	}

	res := 0
	steps := make([]int, 0)

	_, s := reduce(input, steps, true)
	for i := len(s) - 1; i >= 0; i-- {
		res = s[i] - res
	}

	return res
}

// reduce slice to zeros
func reduce(input []int, steps []int, isSecondPuzzle bool) ([]int, []int) {
	if onlyZeros(input) {
		return input, steps
	}

	res := make([]int, 0)
	for i := 0; i < len(input)-1; i++ {
		res = append(res, input[i+1]-input[i])
	}

	n := res[(len(res) - 1)]
	if isSecondPuzzle {
		n = input[0]
	}
	steps = append(steps, n)

	r, s := reduce(res, steps, isSecondPuzzle)

	return r, s
}

// check if slice contains only zeros
func onlyZeros(input []int) bool {
	res := true
	for i := 0; i < len(input); i++ {
		res = res && (input[i] == 0)
	}
	return res
}

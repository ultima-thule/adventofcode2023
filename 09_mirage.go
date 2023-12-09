package main

// Solve puzzle no 1
func puzzle09_1(input []int) int {
	if input == nil {
		return 0
	}

	var res int = input[len(input)-1]
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

	var res int = 0
	steps := make([]int, 0)

	_, n := reduce(input, steps, true)
	for i := len(n) - 1; i >= 0; i-- {
		res = n[i] - res
	}

	return res
}

// reduce slice to zeros
func reduce(input []int, steps []int, isLeft bool) ([]int, []int) {
	if checkZeros(input) {
		return input, steps
	}

	res := make([]int, 0)
	for i := 0; i < len(input)-1; i++ {
		res = append(res, input[i+1]-input[i])
	}

	step := res[(len(res) - 1)]
	if isLeft {
		step = input[0]
	}
	steps = append(steps, step)

	r, s := reduce(res, steps, isLeft)

	return r, s
}

// check if slice contains only zeros
func checkZeros(input []int) bool {
	res := true
	for k := 0; k < len(input); k++ {
		res = res && (input[k] == 0)
	}
	return res
}

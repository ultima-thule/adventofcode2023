package main

import (
	"bufio"
	"fmt"
)

// Day 08 solution
func mirage(filename string, calcFun func([]int) int, prepFun func(string) []int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	var result int

	for fileScanner.Scan() {
		text := fileScanner.Text()
		result += calcFun(prepFun(text))
	}

	return result
}

// Solve puzzle no 1
func puzzle09_1(input []int) int {
	if input == nil {
		return 0
	}

	fmt.Println(input)

	var res int = input[len(input)-1]
	steps := make([]int, 0)

	p, n := reduce(input, steps)
	for i := 0; i < len(n); i++ {
		res += n[i]
	}
	fmt.Println("reduced: ", p)
	fmt.Println("steps:", n)
	fmt.Println("res:", res)

	return res
}

func puzzle09_2(input []int) int {
	if input == nil {
		return 0
	}

	fmt.Println(input)

	var res int = 0
	steps := make([]int, 0)

	_, n := reduce2(input, steps)
	// n = append(n, 0)
	for i := len(n) - 1; i >= 0; i-- {
		// fmt.Println("Back: step - x = res ", n[i], " - ", res, " = ", n[i]-res)
		res = n[i] - res
	}
	// fmt.Println("reduced 2: ", p)
	// fmt.Println("steps 2:", n)
	// fmt.Println("res 2:", res)
	// fmt.Println(" ")

	return res
}

func reduce(input []int, steps []int) ([]int, []int) {
	if checkZeros(input) {
		return input, steps
	}

	res := make([]int, 0)
	for i := 0; i < len(input)-1; i++ {
		res = append(res, input[i+1]-input[i])
	}
	steps = append(steps, res[(len(res)-1)])
	p, n := reduce(res, steps)

	return p, n
}

func reduce2(input []int, steps []int) ([]int, []int) {
	if checkZeros(input) {
		return input, steps
	}

	res := make([]int, 0)
	for i := 0; i < len(input)-1; i++ {
		res = append(res, input[i+1]-input[i])
	}
	steps = append(steps, input[0])
	p, n := reduce2(res, steps)

	return p, n
}

// check if slice contains only zeros
func checkZeros(input []int) bool {
	res := true
	for k := 0; k < len(input); k++ {
		res = res && (input[k] == 0)
	}
	return res
}

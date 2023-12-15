package main

import (
	"strconv"
	"strings"
)

// Solve puzzle no 1
func puzzle13(inputRow []int64, inputCol []int64) int {
	res := 0
	resReflRows := findMirrorLines(inputRow, true)
	resReflCols := findMirrorLines(inputCol, false)

	res = calcResult(resReflRows, true) + calcResult(resReflCols, false)

	return res
}

// Solve puzzle no 2
func puzzle13_2(inputRow []int64, inputCol []int64) int {
	res := 0
	origReflRows := findMirrorLines(inputRow, true)
	origReflCols := findMirrorLines(inputCol, false)

	resReflRows := findMirrorLinesWithError(inputRow, true, origReflRows)
	resReflCols := findMirrorLinesWithError(inputCol, false, origReflCols)

	res = calcResult(resReflRows, true) + calcResult(resReflCols, false)

	return res
}

// convert line to binary representation
func convertToInt(input []string) []int64 {
	ret := make([]int64, len(input))

	for k, v := range input {
		text := v
		tmp := strings.ReplaceAll(text, ".", "0")
		tmp = strings.ReplaceAll(tmp, "#", "1")
		num, _ := strconv.ParseInt(tmp, 2, 64)
		ret[k] = num
	}

	return ret
}

// builid matrix and its transpose in binary representations
func prepData13(input []string) ([]int64, []int64) {
	retRow := convertToInt(input)
	cols := transposeMatrix(input)
	retCol := convertToInt(cols)

	return retRow, retCol
}

// find normal reflection lines (puzzle 1)
func findMirrorLines(input []int64, isRow bool) []int {
	res := make([]int, 0)
	refl := findReflection(input)
	if len(refl) > 0 {
		for i := 0; i < len(refl); i++ {
			bFound := checkReflectionLine(refl[i], input)
			if bFound {
				res = append(res, refl[i])
			}
		}
	}
	return res
}

// find reflection lines with errors (puzzle 2)
func findMirrorLinesWithError(input []int64, isRow bool, reject []int) []int {
	res := make([]int, 0)
	refl := findAllReflections(input)
	if len(refl) > 0 {
		for i := 0; i < len(refl); i++ {
			bFound := checkReflectionLineWithError(refl[i], input)
			if bFound && !isInOld(reject, refl[i]) {
				res = append(res, refl[i])
			}
		}
	}
	return res
}

// check if found solution from puzzle 2 is one of the solutions from puzzle 1
func isInOld(reject []int, found int) bool {
	for i := 0; i < len(reject); i++ {
		if reject[i] == found {
			return true
		}
	}
	return false
}

// calculate final result
func calcResult(results []int, isRow bool) int {
	res := 0
	for i := 0; i < len(results); i++ {
		if isRow {
			res += 100 * (results[i] + 1)
		} else {
			res += (results[i] + 1)
		}
	}
	return res
}

// find reflection => same neighbouring numbers
func findReflection(input []int64) []int {
	res := make([]int, 0)

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			res = append(res, i)
		}
	}

	return res
}

// return all possible reflections
func findAllReflections(input []int64) []int {
	res := make([]int, 0)

	for i := 0; i < len(input); i++ {
		res = append(res, i)
	}
	return res
}

// check if it is a reflection
func checkReflectionLine(reflRow int, input []int64) bool {
	ret := true

	leftLen := reflRow
	rigthLen := len(input) - 1 - reflRow - 1

	cnt := min(leftLen, rigthLen)

	for i := 0; i <= cnt; i++ {
		ret = ret && input[reflRow-i] == input[reflRow+1+i]
	}

	return ret
}

// check if there is a reflection wih exactly one error
func checkReflectionLineWithError(reflRow int, input []int64) bool {
	leftLen := reflRow
	rigthLen := len(input) - 1 - reflRow - 1

	ret := -1
	cycles := -1

	for i := 0; i <= min(leftLen, rigthLen); i++ {
		cycles = i
		if IsXorPowerOfTwo(input[reflRow-i], input[reflRow+1+i]) {
			ret += 1
		}
	}
	return cycles != -1 && (cycles == ret)
}

// check if result is a power of two
func IsXorPowerOfTwo(x1 int64, x2 int64) bool {
	xor := x1 ^ x2
	return (xor & (xor - 1)) == 0
}

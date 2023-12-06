package main

import (
	"strings"
)

// Day 04 solution
func waitForIt(filename string, calcFun func([]string) int, prepFun func(data []string) []string) int {
	input := readInput(filename, calcFun, prepFun)

	return calcFun(input)
}

// Solve puzzle no 1 & 2
func race(input []string) int {
	if input == nil || len(input) < 2 {
		return 0
	}

	var res int = 1

	time := splitToInts(input[0])
	dist := splitToInts(input[1])

	racesCnt := len(time)

	for i := 0; i < racesCnt; i++ {
		cntWins := 0
		for j := 0; j < time[i]-1; j++ {
			if travelDistance(j, time[i]) > dist[i] {
				cntWins++
			}
		}
		res *= cntWins
	}

	return res
}

// calculat travel distance
func travelDistance(hold int, timeMax int) int {
	if hold >= timeMax {
		return 0
	}

	time := timeMax - hold
	dist := hold * time

	return dist
}

// prepare data for puzzle no 2
func prepData(input []string) []string {
	if input == nil || len(input) < 2 {
		return nil
	}

	input[0] = strings.ReplaceAll(input[0], "Time:", "")
	input[0] = strings.ReplaceAll(input[0], " ", "")

	input[1] = strings.ReplaceAll(input[1], "Distance:", "")
	input[1] = strings.ReplaceAll(input[1], " ", "")

	return input
}

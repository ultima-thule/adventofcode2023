package main

import (
	"bufio"
	"fmt"
)

// Day 04 solution
func waitForIt(filename string, calcFun func([]string) int) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	var result int
	input := []string{}

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	result = calcFun(input)

	return result
}

// Solve puzzle no 1
func race(input []string) int {
	var res int = 1

	// translated := false
	time := splitToInts(input[0])
	dist := splitToInts(input[1])

	// fmt.Println("Time: ", time)
	// fmt.Println("Distance: ", dist)

	racesCnt := len(time)

	for i := 0; i < racesCnt; i++ {
		// fmt.Println("Race: ", i, " time ", time[i], " distance ", dist[i])
		cntWins := 0
		for j := 0; j < time[i]-1; j++ {
			// fmt.Println("Check ", j)
			if travelDistance(j, time[i]) > dist[i] {
				cntWins++
			}
		}
		fmt.Println("no of wins: ", cntWins)
		res *= cntWins
	}

	return res
}

func travelDistance(hold int, timeMax int) int {
	if hold >= timeMax {
		return 0
	}

	time := timeMax - hold
	dist := hold * time

	return dist
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "03", "day of AoC competition in format dd with trailing zero")
	flag.Parse()

	testDataOne := "day" + day + "/input_test1.txt"
	realDataOne := "day" + day + "/input1.txt"
	testDataTwo := "day" + day + "/input_test2.txt"
	realDataTwo := "day" + day + "/input2.txt"

	switch day {
	case "01":
		fmt.Println(trebuchet(testDataOne, 1))
		fmt.Println(trebuchet(realDataOne, 1))
		fmt.Println(trebuchet(testDataTwo, 2))
		fmt.Println(trebuchet(realDataTwo, 2))
	case "02":
		fmt.Println(cubeConudrum(testDataOne, calcSum))
		fmt.Println(cubeConudrum(realDataOne, calcSum))
		fmt.Println(cubeConudrum(testDataTwo, calcPower))
		fmt.Println(cubeConudrum(realDataTwo, calcPower))
	case "03":
		fmt.Println(gearRatios(testDataOne, findParts))
		fmt.Println(gearRatios(realDataOne, findParts))
		fmt.Println(gearRatios(testDataTwo, findAdjacentGears))
		fmt.Println(gearRatios(realDataTwo, findAdjacentGears))
	case "04":
		fmt.Println(scratchcards(testDataOne, countWorth))
		fmt.Println(scratchcards(realDataOne, countWorth))
		fmt.Println(scratchcards(testDataTwo, countTotalCards))
		fmt.Println(scratchcards(realDataTwo, countTotalCards))
	case "05":
		// fmt.Println(fertilizer(testDataOne, puzzle1))
		// fmt.Println(fertilizer(realDataOne, puzzle1))
		fmt.Println(fertilizer(testDataTwo, puzzle2))
		// fmt.Println(fertilizer(realDataTwo, puzzle2))
	case "06":
		fmt.Println((waitForIt(testDataOne, race, nil)))
		fmt.Println((waitForIt(realDataOne, race, nil)))
		fmt.Println(waitForIt(testDataTwo, race, prepData))
		fmt.Println(waitForIt(realDataTwo, race, prepData))
	}
}

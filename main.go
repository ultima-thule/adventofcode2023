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
		fmt.Println(trebuchet(testDataOne, 1), "\n")
		fmt.Println(trebuchet(realDataOne, 1), "\n")
		fmt.Println(trebuchet(testDataTwo, 2), "\n")
		fmt.Println(trebuchet(realDataTwo, 2), "\n")
	case "02":
		fmt.Println(cubeConudrum(testDataOne, calcSum), "\n")
		fmt.Println(cubeConudrum(realDataOne, calcSum), "\n")
		fmt.Println(cubeConudrum(testDataTwo, calcPower), "\n")
		fmt.Println(cubeConudrum(realDataTwo, calcPower), "\n")
	case "03":
		fmt.Println(gearRatios(testDataOne, findParts), "\n")
		fmt.Println(gearRatios(realDataOne, findParts), "\n")
		fmt.Println(gearRatios(testDataTwo, findAdjacentGears), "\n")
		fmt.Println(gearRatios(realDataTwo, findAdjacentGears), "\n")
	case "04":
		fmt.Println(scratchcards(testDataOne, countWorth), "\n")
		fmt.Println(scratchcards(realDataOne, countWorth), "\n")
		fmt.Println(scratchcards(testDataTwo, countTotalCards), "\n")
		fmt.Println(scratchcards(realDataTwo, countTotalCards), "\n")
	case "05":
		fmt.Println(fertilizer(testDataOne, puzzle1, nil), "\n")
		fmt.Println(fertilizer(realDataOne, puzzle1, nil), "\n")
		fmt.Println(fertilizer(testDataTwo, puzzle2, nil), "\n")
		fmt.Println(fertilizer(realDataTwo, puzzle2, nil), "\n")
	case "06":
		fmt.Println((waitForIt(testDataOne, race, nil)), "\n")
		fmt.Println((waitForIt(realDataOne, race, nil)), "\n")
		fmt.Println(waitForIt(testDataTwo, race, prepData), "\n")
		fmt.Println(waitForIt(realDataTwo, race, prepData), "\n")
	case "07":
		fmt.Println((camelCards(testDataOne, playGame, prepData07, false)), "\n")
		fmt.Println((camelCards(realDataOne, playGame, prepData07, false)), "\n")
		fmt.Println(camelCards(testDataTwo, playGame, prepData07, true), "\n")
		fmt.Println(camelCards(realDataTwo, playGame, prepData07, true), "\n")
	case "08":
		fmt.Println((hauntedWasteland(testDataOne, travel1, prepData08)), "\n")
		fmt.Println((hauntedWasteland(realDataOne, travel1, prepData08)), "\n")
		fmt.Println(hauntedWasteland(testDataTwo, travel2, prepData08), "\n")
		fmt.Println(hauntedWasteland(realDataTwo, travel2, prepData08), "\n")

	}
}

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
	// testDataThree := "day" + day + "/input_test3.txt"

	switch day {
	case "01":
		fmt.Println(trebuchet(testDataOne, 1))
		fmt.Println(trebuchet(realDataOne, 1))
		fmt.Println(trebuchet(testDataTwo, 2))
		fmt.Println(trebuchet(realDataTwo, 2))
	case "02":
		fmt.Println(cubeConudrum(testDataOne, puzzle02_1))
		fmt.Println(cubeConudrum(realDataOne, puzzle02_1))
		fmt.Println(cubeConudrum(testDataTwo, puzzle02_2))
		fmt.Println(cubeConudrum(realDataTwo, puzzle02_2))
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
		fmt.Println(fertilizer(testDataOne, puzzle05_1, nil))
		fmt.Println(fertilizer(realDataOne, puzzle05_1, nil))
		fmt.Println(fertilizer(testDataTwo, puzzle05_2, nil))
		fmt.Println(fertilizer(realDataTwo, puzzle05_2, nil))
	case "06":
		fmt.Println(waitForIt(testDataOne, puzzle06, nil))
		fmt.Println(waitForIt(realDataOne, puzzle06, nil))
		fmt.Println(waitForIt(testDataTwo, puzzle06, prepData))
		fmt.Println(waitForIt(realDataTwo, puzzle06, prepData))
	case "07":
		fmt.Println(camelCards(testDataOne, puzzle07, prepData07, false))
		fmt.Println(camelCards(realDataOne, puzzle07, prepData07, false))
		fmt.Println(camelCards(testDataTwo, puzzle07, prepData07, true))
		fmt.Println(camelCards(realDataTwo, puzzle07, prepData07, true))
	case "08":
		fmt.Println(hauntedWasteland(testDataOne, puzzle08_1, prepData08))
		fmt.Println(hauntedWasteland(realDataOne, puzzle08_1, prepData08))
		fmt.Println(hauntedWasteland(testDataTwo, puzzle08_2, prepData08))
		fmt.Println(hauntedWasteland(realDataTwo, puzzle08_2, prepData08))
	case "09":
		fmt.Println(mirage(testDataOne, puzzle09_1, prepData09))
		fmt.Println(mirage(realDataOne, puzzle09_1, prepData09))
		fmt.Println(mirage(testDataTwo, puzzle09_2, prepData09))
		fmt.Println(mirage(realDataTwo, puzzle09_2, prepData09))
	case "10":
		// fmt.Println(pipeMaze(testDataOne, puzzle10_1))
		// fmt.Println(pipeMaze(testDataTwo, puzzle10_1))
		// fmt.Println(pipeMaze(realDataOne, puzzle10_1))
		fmt.Println(pipeMaze(testDataOne, puzzle10_2))
		fmt.Println(pipeMaze(testDataTwo, puzzle10_2))
		// fmt.Println(pipeMaze(testDataThree, puzzle10_2))
		fmt.Println(pipeMaze(realDataTwo, puzzle10_2))

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		matchStr(fileScanner.Text())
	}
	return 0
}

func matchStr(text string) {
	pattern := re.MustCompile("(\\d+) green")
	// for m in re.FindStringSubmatch(text) {

	//}

	split := pattern.Split(text, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}
}

func main() {
	fmt.Println("---- Test DataSet1 ----")
	fmt.Println(readFile("input_test1.txt"))

	// fmt.Println("---- DataSet 1 ----")
	// fmt.Println(readFile("input1.txt", 1))

	// fmt.Println("---- Test DataSet 2 ----")
	// fmt.Println(readFile("input_test2.txt", 2))

	// fmt.Println("---- DataSet 2 ----")
	// fmt.Println(readFile("input2.txt", 2))
}

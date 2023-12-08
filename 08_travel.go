package main

// Day 08 solution
func hauntedWasteland(filename string, calcFun func(map[string][]string, string) int, prepFun func(data []string) map[string][]string) int {
	input, path := readInput08(filename, prepFun)
	return calcFun(input, path)
}

// Solve puzzle no 1 && 2
func travel(input map[string][]string, path string) int {
	if input == nil || path == "" {
		return 0
	}

	var res int = 0
	lastNode := "AAA"
	// fmt.Println("Path:", path)

	for {
		found, node, cnt := goThroughPath(input, path, lastNode)
		// fmt.Println("Found: ", found, " node: ", node, " cnt ", cnt, " res ", res)
		res += cnt
		lastNode = node
		if found == true {
			break
		}
	}

	return res
}

func goThroughPath(input map[string][]string, path string, lastNode string) (bool, string, int) {
	var found bool = false
	var cnt int = 0

	for _, c := range path {
		// fmt.Println("Let:", string(c))
		if string(c) == "L" {
			lastNode = input[lastNode][0]
		}
		if string(c) == "R" {
			lastNode = input[lastNode][1]
		}
		cnt++
		if lastNode == "ZZZ" {
			found = true
			break
		}
	}
	return found, lastNode, cnt
}

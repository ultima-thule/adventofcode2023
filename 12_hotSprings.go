package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Spring struct {
	condition string
	groups    []int
	pattern   int64
	blocksize int
	queMarks  int64
	regex     string
}

func hotSprings(filename string, ver2 bool) int {
	f := readFile(filename)
	defer closeFile(f)

	fileScanner := bufio.NewScanner(f)

	res := 0

	for fileScanner.Scan() {
		s := prepData12(fileScanner.Text(), ver2)
		res += puzzle12(s)
	}

	return res
}

func prepData12(input string, ver2 bool) Spring {
	ret := Spring{condition: "", groups: make([]int, 0)}

	splitted := strings.Split(input, " ")

	if ver2 {
		tmp1 := splitted[0]
		tmp2 := splitted[1]
		fmt.Println("Here:", tmp1, tmp2)
		for i := 0; i < 4; i++ {
			tmp1 += fmt.Sprintf("?%s", splitted[0])
			tmp2 += fmt.Sprintf(",%s", splitted[1])
		}
		splitted[0] = tmp1
		splitted[1] = tmp2
	}

	s := splitted[0]
	ret.condition = s
	s = strings.ReplaceAll(s, "#", "1")
	s = strings.ReplaceAll(s, ".", "0")
	s = strings.ReplaceAll(s, "?", "0")
	ret.pattern, _ = strconv.ParseInt(s, 2, 64)
	s = ret.condition
	s = strings.ReplaceAll(s, "?", "1")
	ret.queMarks, _ = strconv.ParseInt(s, 2, 64)

	var tmp []int
	tmp, ret.blocksize = splitToIntsComa(splitted[1])

	for _, v := range splitted[0] {
		if v == '.' {
			ret.regex += "0"
			continue
		}
		if v == '#' {
			ret.regex += "1"
			continue
		}
		if v == '?' {
			ret.regex += "[0-1]"
			continue
		}
	}

	ret.groups = make([]int, 0)
	for i := len(tmp) - 1; i >= 0; i-- {
		ret.groups = append(ret.groups, tmp[i])
	}

	return ret
}

// Solve puzzle no 1
func puzzle12(input Spring) int {
	solution := make([]int64, 0)

	spaces := len(input.condition) - input.blocksize - len(input.groups) + 1
	calcPerm(input, 0, spaces, 0, 0, &solution)

	return len(solution)
}

func calcPerm(input Spring, idx int, spaces int, perm int64, shift int, sol *[]int64) {
	if idx == len(input.groups) {
		if input.pattern&perm == input.pattern {
			binVal := fmt.Sprintf("%b", perm)
			toCheck := fmt.Sprintf("%0*s", len(input.condition), binVal)

			pt := regexp.MustCompile(input.regex)
			if pt.Match([]byte(toCheck)) {
				*sol = append(*sol, perm)
			}
		}
		return
	}

	for it := spaces; it >= 0; it-- {
		b := bits(input.groups[idx])
		calcPerm(input, idx+1, it, perm|(b<<shift), shift+input.groups[idx]+1, sol)
		shift++
	}
}

func bits(b int) int64 {
	return (1 << b) - 1
}

func splitToIntsComa(text string) ([]int, int) {
	tmp := strings.Split(text, ",")
	values := make([]int, 0, len(tmp))
	sumOfBlocks := 0

	for _, raw := range tmp {
		// fmt.Println(raw)
		v, err := strconv.Atoi(raw)
		if err != nil {
			// log.Print(err)
			continue
		}
		values = append(values, v)
		sumOfBlocks += v
	}
	return values, sumOfBlocks
}

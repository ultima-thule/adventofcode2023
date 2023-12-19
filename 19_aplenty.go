package main

import (
	"github.com/Knetic/govaluate"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Workflow struct {
	src string
	dst string
}

// Solve puzzle no 18 part 1 & 2
func puzzle19(workflows map[string][]Workflow, parts []Part) int {
	res := 0

	for _, v := range parts {
		res += goWorkflow(v, workflows, "in")
	}

	return res
}

func goWorkflow(p Part, wf map[string][]Workflow, start string) int {
	res := 0
	for _, v := range wf[start] {
		if v.src != "*" {

			expr, _ := govaluate.NewEvaluableExpression(v.src)

			params := make(map[string]interface{}, 4)
			params["x"] = p.x
			params["m"] = p.m
			params["a"] = p.a
			params["s"] = p.s

			result, _ := expr.Evaluate(params)
			if result == true {
				if v.dst == "A" {
					return getRating(p)
				} else if v.dst == "R" {
					return 0
				} else {
					return goWorkflow(p, wf, v.dst)
				}
			}
		} else {
			if v.dst == "A" {
				return getRating(p)
			} else if v.dst == "R" {
				return 0
			} else {
				return goWorkflow(p, wf, v.dst)
			}
		}
	}

	return res
}

func getRating(p Part) int {
	return p.x + p.m + p.a + p.s
}

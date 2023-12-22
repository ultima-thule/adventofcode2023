package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type PartRange struct {
	xFrom int
	xTo   int
	mFrom int
	mTo   int
	aFrom int
	aTo   int
	sFrom int
	sTo   int
	dst   string
}

type Workflow struct {
	src         string
	dst         string
	src_var     string
	src_greater bool
	src_num     int
}

func (pr PartRange) String() string {
	return fmt.Sprintf("%d_%d_%d_%d_%d_%d_%d_%d", pr.xFrom, pr.xTo, pr.mFrom, pr.mTo, pr.aFrom, pr.aTo, pr.sFrom, pr.sTo)
}

// Solve puzzle no 19 part 1
func puzzle19(workflows map[string][]Workflow, parts []Part) int {
	res := 0

	for _, v := range parts {
		res += goWorkflow(v, workflows, "in")
	}

	return res
}

// Solve puzzle no 19 part 2
func puzzle19_2(workflows map[string][]Workflow, parts []Part) int {
	todo := make(map[string]PartRange, 0)
	completed := make(map[string]PartRange, 0)

	pr := PartRange{xFrom: 1, xTo: 4000, mFrom: 1, mTo: 4000, aFrom: 1, aTo: 4000, sFrom: 1, sTo: 4000, dst: "in"}
	todo[pr.String()] = pr

	goWorkflowsRange(todo, completed, workflows)

	sum := 0
	for _, v := range completed {
		if v.dst == "A" {
			xCnt := v.xTo - v.xFrom + 1
			mCnt := v.mTo - v.mFrom + 1
			aCnt := v.aTo - v.aFrom + 1
			sCnt := v.sTo - v.sFrom + 1
			sum += xCnt * mCnt * aCnt * sCnt
		}
	}
	return sum
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

func goWorkflowsRange(todo map[string]PartRange, completed map[string]PartRange, wf map[string][]Workflow) {
	if len(todo) == 0 {
		return
	}

	// get first elem of map
	for key := range todo {
		current := todo[key]
		dst := current.dst

		// apply all workflows
		for _, w := range wf[dst] {
			p, greater := divideRanges(w, current)
			if greater == true {
				// right range to completed list if it A or R
				// if not add to open
				if p[1].dst == "A" || p[1].dst == "R" {
					completed[p[1].String()] = p[1]
				} else {
					todo[p[1].String()] = p[1]
				}
				// set left range as current
				if len(p) == 1 {
					break
				}
				current = p[0]
			} else {
				// left range to completed list if it A or R
				// if not add to open
				if p[0].dst == "A" || p[0].dst == "R" {
					completed[p[0].String()] = p[0]
				} else {
					todo[p[0].String()] = p[0]
				}
				// set right range as current
				if len(p) == 1 {
					break
				}
				current = p[1]
			}
		}
		// move from todo to completed
		completed[key] = todo[key]
		delete(todo, key)
		break
	}

	goWorkflowsRange(todo, completed, wf)
}

func getRating(p Part) int {
	return p.x + p.m + p.a + p.s
}

func divideRanges(w Workflow, r PartRange) ([]PartRange, bool) {
	ret := []PartRange{}

	p1 := PartRange{xFrom: r.xFrom, xTo: r.xTo, mFrom: r.mFrom, mTo: r.mTo, aFrom: r.aFrom, aTo: r.aTo, sFrom: r.sFrom, sTo: r.sTo}
	p2 := PartRange{xFrom: r.xFrom, xTo: r.xTo, mFrom: r.mFrom, mTo: r.mTo, aFrom: r.aFrom, aTo: r.aTo, sFrom: r.sFrom, sTo: r.sTo}
	switch w.src_var {
	case "x":
		if w.src_greater == false {
			if r.xFrom < w.src_num-1 {
				p1.xTo = w.src_num - 1
				p2.xFrom = w.src_num
				p1.dst = w.dst
			}
		} else {
			if r.xTo > w.src_num {
				p1.xTo = w.src_num
				p2.xFrom = w.src_num + 1
				p2.dst = w.dst
			}
		}
	case "m":
		if w.src_greater == false {
			if r.mFrom < w.src_num-1 {
				p1.mTo = w.src_num - 1
				p2.mFrom = w.src_num
				p1.dst = w.dst
			}
		} else {
			if r.mTo > w.src_num {
				p1.mTo = w.src_num
				p2.mFrom = w.src_num + 1
				p2.dst = w.dst
			}
		}
	case "a":
		if w.src_greater == false {
			if r.aFrom < w.src_num-1 {
				p1.aTo = w.src_num - 1
				p2.aFrom = w.src_num
				p1.dst = w.dst
			}
		} else {
			if r.aTo > w.src_num {
				p1.aTo = w.src_num
				p2.aFrom = w.src_num + 1
				p2.dst = w.dst
			}
		}
	case "s":
		if w.src_greater == false {
			if r.sFrom < w.src_num-1 {
				p1.sTo = w.src_num - 1
				p2.sFrom = w.src_num
				p1.dst = w.dst
			}
		} else {
			if r.sTo > w.src_num {
				p1.sTo = w.src_num
				p2.sFrom = w.src_num + 1
				p2.dst = w.dst
			}
		}
	case "*":
		p1.dst = w.dst
		return append(ret, p1), w.src_greater
	}
	if w.src_greater == true {
		ret = append(ret, p1)
		ret = append(ret, p2)
	} else {
		ret = append(ret, p1)
		ret = append(ret, p2)
	}

	return ret, w.src_greater
}

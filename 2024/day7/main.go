package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
)

// Let's solve this with recursion

type equation struct {
	terms []int
	lenn  int
	sol   int
}

func (eq equation) String() string {
	return fmt.Sprintf("%d -> %v", eq.sol, eq.terms)
}

func parseProblem(input []string) (equations []equation) {
	for _, line := range input {
		e := equation{}
		terms := strings.Split(line, ":")

		sol, err := strconv.Atoi(terms[0])
		utils.DieOnError(err)
		e.sol = sol

		e.terms = lo.Map(strings.Split(strings.TrimSpace(terms[1]), " "), func(s string, _ int) int {
			res, err := strconv.Atoi(s)
			utils.DieOnError(err)
			return res
		})
		e.lenn = len(e.terms)
		equations = append(equations, e)
	}
	return
}

func (eq *equation) recursiveCheck(index int, tot int, concatenator bool) (res bool) {
	if ((index) == eq.lenn) && tot == eq.sol {
		return true
	}
	if index < eq.lenn {
		res = eq.recursiveCheck(index+1, tot*eq.terms[index], concatenator)
		if !res {
			res = eq.recursiveCheck(index+1, tot+eq.terms[index], concatenator)
		}
		if !res && concatenator {
			s := fmt.Sprintf("%d%d", tot, eq.terms[index])
			concatenated, err := strconv.Atoi(s)
			utils.DieOnError(err)
			res = eq.recursiveCheck(index+1, concatenated, concatenator)
		}

	}

	return res
}

func StarOne(input []string, concatenator bool) {
	start := time.Now()
	res := 0
	equations := parseProblem(input)
	for _, eq := range equations {
		if ok := eq.recursiveCheck(1, eq.terms[0], concatenator); ok {
			res += eq.sol
		}
	}

	if !concatenator {
		fmt.Printf("Star One: %d, in %fs\n", res, time.Now().Sub(start).Seconds())
	} else {
		fmt.Printf("Star Two: %d, in %fs\n", res, time.Now().Sub(start).Seconds())
	}
}

func StarTwo(input []string) {
	StarOne(input, true)
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	StarOne(input, false)
	StarTwo(input)
}

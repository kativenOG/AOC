package main

import (
	"flag"
	"fmt"
	"github.com/AOC/2024/utils"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

// DIRECTIONS:
const (
	UNDEFINED = iota
	INCREASING
	DECREASING
)

func intAbs(a, b int) (res int) {
	res = a - b
	if res < 0 {
		res = -res
	}
	return
}

func parseDir(previous, current, currentDirection int) int {
	if (currentDirection == UNDEFINED) && (previous == current) {
		return UNDEFINED
	} else if previous > current {
		return DECREASING
	}
	return INCREASING
}

func parseReport(levels []int) bool {
	var direction int = UNDEFINED
	previous := levels[0]
	for _, level := range levels[1:] {
		parsedDirection := parseDir(previous, level, direction)
		if direction == 0 {
			direction = parsedDirection
		}
		if diff := intAbs(previous, level); (diff < 1 || diff > 3) ||
			(parsedDirection != direction) {
			return false
		}
		previous = level
	}
	return true
}

// parseReportDampened if the first check fails it checks iteritavely for a functioning levels set
func parseReportDampened(levels []int) (res bool) {
	res = parseReport(levels)
	if res {
		return
	}

	for index, _ := range levels {
		dampenedLevels := make([]int, len(levels))
		copy(dampenedLevels, levels)
		dampenedLevels = append(dampenedLevels[:index], dampenedLevels[index+1:]...)

		if res = parseReport(dampenedLevels); res {
			break
		}
	}

	return
}

func starOne(inputs []string, dampened bool) {
	var validReports int
	for _, line := range inputs {
		stringLevels := strings.Split(line, " ")
		levels := lo.Map(stringLevels, func(val string, _ int) int {
			res, err := strconv.Atoi(strings.TrimSpace(val))
			utils.DieOnError(err)
			return res
		})
		var ok bool
		if dampened {
			ok = parseReportDampened(levels)
		} else {
			ok = parseReport(levels)
		}
		if ok {
			validReports += 1
		}
	}

	if dampened {
		fmt.Printf("Star Two: %d\n", validReports)
		return
	}
	fmt.Printf("Star One: %d\n", validReports)
}

func starTwo(inputs []string) {
	starOne(inputs, true)
}

func main() {
	filename, _ := utils.ParseFlags()
	input := utils.ParseInputFile(filename)

	starOne(input, false)
	starTwo(input)

}
